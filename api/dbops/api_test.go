/*
 * @Author: gongluck
 * @Date: 2020-05-22 15:29:57
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-05-22 16:14:35
 */

package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUser)
	t.Run("Get", TestGetUser)
	t.Run("Del", TestDelUser)
	t.Run("Reget", TestRegetUser)
}

func TestAddUser(t *testing.T) {
	err := AddUserCredential("gongluck", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	pwd, err := GetUserCredential("gongluck")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func TestDelUser(t *testing.T) {
	err := DeleteUser("gongluck", "123")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func TestRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("gongluck")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Error of DelUser.")
	}
}
