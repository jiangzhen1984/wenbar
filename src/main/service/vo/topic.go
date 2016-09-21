

package vo


import (

)



type Topic struct {

     Id        uint64
     Title     string 
     Price     float32
     Count     uint32

     AnsUser   *User
}
