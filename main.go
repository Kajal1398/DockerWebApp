package main
import(
	"net/http"
	"text/template"
)
var tpl *template.Template
type Todo struct {
    Name string
	Recommanded_For string
	Maintenance_level string
	Lifespan string
	Temperament string
	Health_risk string
	Image_url string
    Url  string
}



type TodoPageData struct {
    Todos     []Todo
}
func init(){
	tpl=template.Must(template.ParseFiles("form.html","animal.html","information.html"))
}
func main(){

	http.HandleFunc("/",indexfunc)
	http.HandleFunc("/animal",animal_list)
	http.HandleFunc("/info",animal_info)
	http.ListenAndServe(":8080",nil)
	
}
func indexfunc(w http.ResponseWriter,r *http.Request){
	tpl.Execute(w,nil)
}
func animal_list(w http.ResponseWriter,r *http.Request){

	data := TodoPageData{
            
            Todos: []Todo{
				{Name: "Airedale Terrier", Recommanded_For : "Small Families", Maintenance_level : "High", Lifespan : "10-14 years", Temperament: "Intelligent", Health_risk : "High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure.",Image_url : "https://images2.minutemediacdn.com/image/upload/c_fill,g_auto,h_1248,w_2220/f_auto,q_auto,w_1100/v1555312680/shape/mentalfloss/istock_79171327_small.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/airedale-terrier/"},
				
				{Name: "Akita", Recommanded_For :"Small families, singles (wary of strangers)", Maintenance_level : "High", Lifespan : "11-15 years", Temperament: "Active", Health_risk :"High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure." ,Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/10/akita-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/akita/"},
				//{Name: "Task 3", Recommanded_For :, Maintenance_level : , Lifespan : , Temperament: , Health_risk : ,Image_url : , Url: "true"},
				{Name: "Alaskan Malamute", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "12-16 years", Temperament: "Intelligent", Health_risk : "High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure.",Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/10/alaskan-malamute-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/alaskan-malamute/"},
				{Name: "American Bulldog", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "8-15 years", Temperament: "friendly", Health_risk :"Higher than average probability of developing health issues during its lifetime, hence the cost to insure is above average." ,Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/11/american-bulldog-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/american-bulldog/"},
				{Name: "American Staffordshire Terrier", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "10-12 years", Temperament: "social", Health_risk : "This breed has an around average probability of having health issues in its lifetime, hence it is one of the more affordable breeds to insure.",Image_url :"https://bowwowinsurance.com.au/wp-content/uploads/2018/10/american-staffordshire-terrier-amstaff-american-staffy-700x700.jpg" , Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/american-staffordshire-terrier/"},
            },
        }
	tpl.ExecuteTemplate(w,"animal.html",data)
}
func animal_info(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"information.html",nil)
}
