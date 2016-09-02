package orm

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

/*
	项目地址: https://github.com/jinzhu/gorm
	项目文档: http://jinzhu.me/gorm/
*/

// import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

func ExampleGorm() {
	db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err := gorm.Open("foundation", "dbname=gorm") // FoundationDB.
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	// You can also use an existing database connection handle
	// dbSql, _ := sql.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, _ := gorm.Open("postgres", dbSql)
	if nil != err {

	}

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// Disable table name's pluralization
	db.SingularTable(true)

	///////////////////////////////Migration///////////////////////////////
	if true {

		// Create table
		db.CreateTable(&User{})

		// Drop table
		db.DropTable(&User{})

		// Automating Migration
		db.AutoMigrate(&User{})
		//		db.AutoMigrate(&User{}, &Product{}, &Order{})
	}

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	///////////////////////////////Create Record///////////////////////////////
	if true {

		db.NewRecord(user) // => returns `true` if primary key is blank

		db.Create(&user)

		db.NewRecord(user) // => return `false` after `user` created

		// Associations will be inserted automatically when save the record
		user := User{
			Name:            "jinzhu",
			BillingAddress:  Address{Address1: "Billing Address - Address 1"},
			ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
			Emails:          []Email{{Email: "jinzhu@example.com"}, {Email: "jinzhu-2@example@example.com"}},
			Languages:       []Language{{Name: "ZH"}, {Name: "EN"}},
		}

		db.Create(&user)
		//// BEGIN TRANSACTION;
		//// INSERT INTO "addresses" (address1) VALUES ("Billing Address - Address 1");
		//// INSERT INTO "addresses" (address1) VALUES ("Shipping Address - Address 1");
		//// INSERT INTO "users" (name,billing_address_id,shipping_address_id) VALUES ("jinzhu", 1, 2);
		//// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu@example.com");
		//// INSERT INTO "emails" (user_id,email) VALUES (111, "jinzhu-2@example.com");
		//// INSERT INTO "languages" ("name") VALUES ('ZH');
		//// INSERT INTO user_languages ("user_id","language_id") VALUES (111, 1);
		//// INSERT INTO "languages" ("name") VALUES ('EN');
		//// INSERT INTO user_languages ("user_id","language_id") VALUES (111, 2);
		//// COMMIT;
	}

	///////////////////////////////Query///////////////////////////////
	if true {
		// Get the first record
		db.First(&user)
		//// SELECT * FROM users ORDER BY id LIMIT 1;

		// Get the last record
		db.Last(&user)
		//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

		// Get all records
		var users []User
		db.Find(&users)
		//// SELECT * FROM users;

		// Get record with primary key
		db.First(&user, 10)
		//// SELECT * FROM users WHERE id = 10;
	}
}

type User struct {
	ID        int
	Birthday  time.Time
	Age       int
	Name      string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
	Num       int    `sql:"AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Emails            []Email       // One-To-Many relationship (has many)
	BillingAddress    Address       // One-To-One relationship (has one)
	BillingAddressID  sql.NullInt64 // Foreign key of BillingAddress
	ShippingAddress   Address       // One-To-One relationship (has one)
	ShippingAddressID int           // Foreign key of ShippingAddress
	IgnoreMe          int           `sql:"-"`                          // Ignore this field
	Languages         []Language    `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	ID         int
	UserID     int    `sql:"index"`                          // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
	Email      string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `sql:"type:varchar(100);unique"`
	Post     sql.NullString `sql:"not null"`
}

type Language struct {
	ID   int
	Name string `sql:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `sql:"index:idx_name_code"` // `unique_index` also works
}
