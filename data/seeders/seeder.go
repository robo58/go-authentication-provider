package seeders

import (
	"github.com/jinzhu/gorm"
	"github.com/robo58/go-authentication-provider/data/models"
)

type Seeder struct {
	Name string
	Run func(*gorm.DB) error
}

func All() []Seeder {
	return []Seeder{
		{
			Name: "CreateAdminRole",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "admin")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateHeadmasterRole",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "headmaster")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateTeacherRole",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "teacher")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateStudentRole",
			Run: func(db *gorm.DB) error {
				err := CreateRole(db, "student")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAdminUser",
			Run: func(db *gorm.DB) error {
				var adminRole models.Role
				err := db.Where("name = ?", "admin").First(&adminRole).Error
				if err != nil {
					return err
				}
				err = CreateUser(db, "Admin", "admin@mail.com", "password")
				if err != nil {
					return err
				}
				var user models.User
				db.Where(&models.User{Name: "Admin"}).First(&user).Association("Roles").Append([]models.Role{adminRole})
				return nil
			},
		},
		{
			Name: "CreateHeadmasterUser",
			Run: func(db *gorm.DB) error {
				var headmasterRole models.Role
				var user models.User
				err := db.Where("name = ?", "headmaster").First(&headmasterRole).Error
				if err != nil {
					return err
				}
				err = CreateUser(db, "Headmaster", "headmaster@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Headmaster"}).First(&user).Association("Roles").Append([]models.Role{headmasterRole})
				return nil
			},
		},
		{
			Name: "CreateTeachers",
			Run: func(db *gorm.DB) error {
				var teacherRole models.Role
				var user1 models.User
				var user2 models.User
				err := db.Where("name = ?", "teacher").First(&teacherRole).Error
				if err != nil {
					return err
				}
				err = CreateUser(db, "Teacher1", "teacher1@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Teacher1"}).First(&user1).Association("Roles").Append([]models.Role{teacherRole})
				err = CreateUser(db, "Teacher2", "teacher2@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Teacher2"}).First(&user2).Association("Roles").Append([]models.Role{teacherRole})
				return nil
			},
		},
		{
			Name: "CreateStudents",
			Run: func(db *gorm.DB) error {
				var studentRole models.Role
				var user1 models.User
				var user2 models.User
				var user3 models.User
				var user4 models.User
				err := db.Where("name = ?", "student").First(&studentRole).Error
				if err != nil {
					return err
				}
				err = CreateUser(db, "Student1", "student1@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Student1"}).First(&user1).Association("Roles").Append([]models.Role{studentRole})
				err = CreateUser(db, "Student2", "student2@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Student2"}).First(&user2).Association("Roles").Append([]models.Role{studentRole})
				err = CreateUser(db, "Student3", "student3@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Student3"}).First(&user3).Association("Roles").Append([]models.Role{studentRole})
				err = CreateUser(db, "Student4", "student3@mail.com", "password")
				if err != nil {
					return err
				}
				db.Where(&models.User{Name: "Student4"}).First(&user4).Association("Roles").Append([]models.Role{studentRole})
				return nil
			},
		},
		{
			Name: "CreateSubjects",
			Run: func(db *gorm.DB) error {
				err := CreateSubject(db, "Matematika", 3)
				if err != nil {
					return err
				}
				err = CreateSubject(db, "Hrvatski", 4)
				if err != nil {
					return err
				}
				err = CreateSubject(db, "Engleski", 4)
				if err != nil {
					return err
				}
				err = CreateSubject(db, "Informatika", 3)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateSchool",
			Run: func(db *gorm.DB) error {
				err := CreateSchool(db, "Testna skola", 2)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateDepartments",
			Run: func(db *gorm.DB) error {
				err := CreateDepartment(db, "1.A", 1, 3)
				if err != nil {
					return err
				}
				err = CreateDepartment(db, "1.B", 1,4)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "AssignSubjectsToDepartments",
			Run: func(db *gorm.DB) error {
				var subjects1 []*models.Subject
				var subjects2 []*models.Subject
				var department1 models.Department
				var department2 models.Department
				db.Where([]int{1,2,3}).Find(&subjects1)
				db.Where([]int{2,3,4}).Find(&subjects2)
				db.First(&department1, 1)
				db.First(&department2, 2)
				department1.Subjects = subjects1
				department2.Subjects = subjects2
				db.Save(&department1)
				db.Save(&department2)
				return nil
			},
		},
		{
			Name: "AssignStudentsToDepartments",
			Run: func(db *gorm.DB) error {
				var students1 []*models.User
				var students2 []*models.User
				var department1 models.Department
				var department2 models.Department
				err := db.Where([]int{5, 6}).Find(&students1).Error
				if err != nil {
					return err
				}
				err = db.Where([]int{7, 8}).Find(&students2).Error
				if err != nil {
					return err
				}
				var departmentStudents1 []*models.DepartmentStudent
				var departmentStudents2 []*models.DepartmentStudent
				for _, student := range students1 {
					departmentStudents1 = append(departmentStudents1, &models.DepartmentStudent{UserId: int(student.ID)})
				}
				for _, student := range students2 {
					departmentStudents2 = append(departmentStudents2, &models.DepartmentStudent{UserId: int(student.ID)})
				}
				err = db.First(&department1, 1).Association("Students").Append(departmentStudents1).Error
				if err != nil {
					return err
				}
				err = db.First(&department2, 2).Association("Students").Append(departmentStudents2).Error
				if err != nil {
					return err
				}

				err = db.Save(&department1).Error
				if err != nil {
					return err
				}
				err=db.Save(&department2).Error
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "AssignSubjectsToStudents",
			Run: func(db *gorm.DB) error {
				var subjects1 []*models.DepartmentSubject
				var subjects2 []*models.DepartmentSubject
				var students1 []*models.DepartmentStudent
				var students2 []*models.DepartmentStudent
				var department1 models.Department
				var department2 models.Department
				db.First(&department1, 1)
				db.First(&department2, 2)
				db.Where(&models.DepartmentSubject{DepartmentId: int(department1.ID)}).Find(&subjects1)
				db.Where(&models.DepartmentSubject{DepartmentId: int(department2.ID)}).Find(&subjects2)
				db.Where(&models.DepartmentStudent{DepartmentId: int(department1.ID)}).Find(&students1)
				db.Where(&models.DepartmentStudent{DepartmentId: int(department2.ID)}).Find(&students2)
				for _, student := range students1 {
					student.Subjects = subjects1
					err := db.Save(&student).Error
					if err != nil {
						return err
					}
				}
				for _, student := range students2 {
					student.Subjects = subjects2
					err := db.Save(&student).Error
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
	}
}