package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

type User struct {
	Profile   Profile `gorm:"ForeignKey:ProfileID;"`
	ProfileID uint
	Email []Email
	Language []Language `gorm:"many2many:user_languages;"`
	gorm.Model
}

type Profile struct {
	Name string
	gorm.Model
}

type Email struct {
	Email   string
	UserID  uint
	User User
	gorm.Model
}

type Language struct {
	gorm.Model
	Name string
}

type TestTable struct {
	MyModel
	Profile   Profile `gorm:"ForeignKey:ProfileID;"`
	ProfileID uint
}

type MyModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedBy uint
	CreatedAt time.Time
	UpdatedBy uint
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func main() {
	automigrate()
	//o2o()
	//o2m()
	//m2m()
	//m2o()
	//updateO2M()
	//updateM2M()
	//relatedAndAssociation()
}

func insertRelation()  {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	db.AutoMigrate(&User{}, &Profile{}, &Email{}, &Language{})

	user := User{
		Profile:Profile{Name:"yyy"},
		Email:[]Email{{Email:"ccc@mail.com"}, {Email:"ddd@mail.com"}},
		Language: []Language{{Name:"TH"}, {Name:"EN"}, {Name:"FR"}},
	}
	db.Create(&user)
}

func updateO2O() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	db.AutoMigrate(&User{}, &Profile{}, &Email{}, &Language{})

	user := User{}
	db.Preload("Profile").First(&user, 14)

	user.Profile.Name = "abc"
	db.Save(&user)

	fmt.Println("endFunc")
}

func updateM2M() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	user := User{}
	db.First(&user, 14)

	l1 := Language{}
	db.First(&l1, 6)

	l2 := Language{}
	db.First(&l2, 7)

	db.Preload("Language").First(&user, 14)
	db.Debug().Model(&user).Association("Language").Replace(&l1, &l2)

	db.Save(&user)

	fmt.Println("endFunc")
}

func o2o() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	user := []User{}
	db.Debug().Preload("Profile").Find(&user)

	for _, u := range user {
		fmt.Println(u.Profile.Name)
	}

	fmt.Println("xxx")
}

func o2m() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	user := []User{}
	db.Debug().Preload("Email").Find(&user)

	for _, u := range user {
		fmt.Println("=====> ",u.ID)
		for _, e := range u.Email {
			fmt.Println(e.Email)
		}
	}

	fmt.Println("xxx")
}

func m2o() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	db.AutoMigrate(&User{}, &Profile{}, &Email{}, &Language{})

	email := []Email{}
	db.Debug().Preload("User").Find(&email)

	for _, e := range email {
		fmt.Println(e.User.ID)
	}

	fmt.Println("xxx")
}

func m2m() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	user := []User{}
	db.Debug().Preload("Language").Find(&user)

	for _, u := range user {
		fmt.Println("=====> ",u.ID)
		for _, e := range u.Language {
			fmt.Println(e.Name)
		}
	}

	fmt.Println("xxx")
}

func relatedAndAssociation() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	user := User{}
	db.Debug().First(&user, 14)

	db.Debug().Model(&user).Related(&user.Profile)
	db.Debug().Model(&user).Related(&user.Email)
	db.Debug().Model(&user).Association("Language").Find(&user.Language)
	db.Debug().Model(&user).Related(&user.Language, "Language")
	db.Debug().Preload("Language").Find(&user)

	fmt.Println("endfunc")
}

func automigrate() {
	db, _ := gorm.Open("mysql", "root:@/testgorm?parseTime=true")
	defer db.Close()

	db.AutoMigrate(&TestTable{})
}