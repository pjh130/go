@echo off


:: 获取目录的名字, 用于start  目录名字.exe
set dirName=
goto   getDirName

:: 删进程文件
taskkill /f /fi "IMAGENAME eq %dirName%.exe" > nul

:: 启动这2个服务以避免exe文件延迟1分钟才能删除 Application Experience和 Program Compatibility Assistant Service
for /f "skip=3 tokens=4" %%i in ('sc query AeLookupSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" (  echo . ) else (  net start "AeLookupSvc" )

for /f "skip=3 tokens=4" %%i in ('sc query PcaSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" (  echo . ) else (  net start "PcaSvc" )


:: 删除旧的exe编译文件
if     exist    %dirName%.exe      (  del %dirName%.exe ) > nul

:: 执行build_go_app子程序：修改gopath环境变量、编译go程序为exe、启动编译的exe程序
goto   build_go_app                :: 不需要进入： %_curDisk%     &    ( cd   %_curPath% )  



:: 获取目录的名字
:getDirName
    set "lj=%~p0"
    set "lj=%lj:\= %"
    for %%a in (%lj%) do set wjj=%%a
    set dirName=%wjj%



:: 编译go程序
:build_go_app

    echo.
    echo 【开始编译go程序为exe可执行文件】：


    :: 执行go编译
    go build -ldflags=" -H windowsgui   "  -gcflags=" -N -l  " 

    echo   已执行编译命令： go build -ldflags=" -H windowsgui   "  -gcflags=" -N -l  " 
    
    :: 删除调试符号：go build -ldflags “-s -w”          -s: 去掉符号信息。  -w: 去掉DWARF调试信息。
    :: 关闭内联优化：go build -gcflags “-N -l”

    echo.
    echo.
    echo.
    
    :: 启动编译的程序.  成功、延时2秒消失； 失败、不消失 暂停   
    if     exist    %dirName%.exe       (    ( choice /t 1 /d y /n >nul )   | ( echo  ok, 编译成功！ 正在打开程序 )  & ( start %dirName%.exe )  )      else    (  echo  编译失败！| pause  )