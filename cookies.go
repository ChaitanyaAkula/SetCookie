package cookies
import (
	"net/http"
	"time"
)
var m,_=time.ParseDuration("325600m")
var expire=time.Now().Add(m)
func SetCookie(w http.ResponseWriter,name,value string){
        http.SetCookie(w,&http.Cookie{
		Name:name,
		Value:value,
		Expires:expire, // persistent cookie , max age = 1 year
		HttpOnly: true,
	})

}
