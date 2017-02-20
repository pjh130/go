package models

import (
	"log"
	"sync"
	"time"

	fs1 "github.com/fsnotify/fsnotify"
	fs2 "github.com/howeyc/fsnotify"
	"github.com/pjh130/go/project/myfile/utils"
)

//文件监控发生的操作
const (
	OpWrite = iota
	OpCreate
	OpRemove
	OpRename
	OpChmod
)

func init() {
	task.files = make(map[string]Status)
	task.Run()
}

var task Task

type Task struct {
	lock  sync.Mutex
	files map[string]Status
}

type Status struct {
	LastTime  time.Time
	Operation int  //监控文件进行的操作
	Working   bool //任务是否进行中
	Fail      int  //失败次数
}

func (t *Task) Add(name string, op int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	old, ok := t.files[name]

	if ok {
		old.Fail = 0
		t.files[name] = old
	} else {
		var v Status
		//		v.LastTime = time.Now()
		v.Working = false
		v.Fail = 0
		v.Operation = op
		t.files[name] = v
	}
}

func (t *Task) Delete(name string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	delete(t.files, name)
}

func (t *Task) Modify(name string, status Status) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.files[name] = status
}

func (t *Task) Fail(name string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	old, ok := t.files[name]
	if ok {
		if old.Fail >= 3 {
			delete(t.files, name)
		} else {
			old.Fail++
			t.files[name] = old
		}
	}
}

func (t *Task) Run() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				t.lock.Lock()
				for key, v := range t.files {
					now := time.Now()
					//1分钟内不会重复处理任务
					if (now.Unix() - v.LastTime.Unix()) < 60 {
						continue
					}

					//任务未进行
					if false == v.Working {
						v.Working = true
						v.LastTime = now
						t.files[key] = v

						if v.Operation == OpWrite {
							go func() {
								//err := FileOverwrite("", key)
								err := FileUpload(key)
								if nil == err {
									t.Delete(key)
								} else {
									//失败了重新处理
									v.Working = false
									v.Fail++
									if v.Fail >= 3 {
										t.Delete(key)
									} else {
										t.Modify(key, v)
									}
								}
							}()
						} else if v.Operation == OpCreate {
							go func() {
								err := FileUpload(key)
								if nil == err {
									t.Delete(key)
								} else {
									//失败了重新处理
									v.Working = false
									v.Fail++
									if v.Fail >= 3 {
										t.Delete(key)
									} else {
										t.Modify(key, v)
									}
								}
							}()
						} else if v.Operation == OpRemove {
							go func() {
								err := FileDelete(key)
								if nil == err {
									t.Delete(key)
								} else {
									//失败了重新处理
									v.Working = false
									v.Fail++
									if v.Fail >= 3 {
										t.Delete(key)
									} else {
										t.Modify(key, v)
									}
								}
							}()
						} else if v.Operation == OpRename {
							go func() {
								err := FileRename(key, "")
								if nil == err {
									t.Delete(key)
								} else {
									//失败了重新处理
									v.Working = false
									v.Fail++
									if v.Fail >= 3 {
										t.Delete(key)
									} else {
										t.Modify(key, v)
									}
								}
							}()
						} else if v.Operation == OpChmod {

						} else {

						}
					} else {
						log.Println("Progress working")
					}
				}
				t.lock.Unlock()
			}
		}
	}()
}

func StartWatch() {
	for _, dir := range utils.Config.Files {
		if len(dir) > 0 {
			go watchDir1(dir)
			//			go watchDir2(dir)
		}
	}
}

func watchDir1(dir string) {
	watcher, err := fs1.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fs1.Write == fs1.Write {
					log.Println("Write file:", event.Name)
					task.Add(event.Name, OpWrite)
				} else if event.Op&fs1.Create == fs1.Create {
					log.Println("Create file:", event.Name)
					task.Add(event.Name, OpCreate)
				} else if event.Op&fs1.Remove == fs1.Remove {
					log.Println("Remove file:", event.Name)
					task.Add(event.Name, OpRemove)
				} else if event.Op&fs1.Rename == fs1.Rename {
					log.Println("Rename file:", event.Name)
					task.Add(event.Name, OpRename)
				} else if event.Op&fs1.Chmod == fs1.Chmod {
					log.Println("Chmod file:", event.Name)
					task.Add(event.Name, OpChmod)
				} else {
					log.Println("Unkown op:")
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	<-done

	return
}

func watchDir2(dir string) {
	watcher, err := fs2.NewWatcher()
	if err != nil {
		log.Println(err)
		return
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev.IsCreate() {
					log.Println("IsCreate", ev.Name)
				}

				if ev.IsDelete() {
					log.Println("IsDelete", ev.Name)
				}

				if ev.IsModify() {
					log.Println("IsModify", ev.Name)
				}

				if ev.IsRename() {
					log.Println("IsRename", ev.Name)
				}

				if ev.IsAttrib() {
					log.Println("IsAttrib", ev.Name)
				}

			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(dir)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Watch path:", dir)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()

}
