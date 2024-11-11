package main

import (
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func main() {
	// ВАРИАНТ 1:  Всё, что вам нужно для подключения профайлера, – импортировать net/http/pprof; необходимые HTTP-обработчики будут зарегистрированы автоматически
	//http.HandleFunc("/", hiHandler)
	//log.Println("Сервер запущен")
	//http.ListenAndServe(":8080", nil)

	// ВАРИАНТ 2: Если ваше веб-приложение использует собственный URL-роутер, необходимо вручную зарегистрировать несколько pprof-адресов:
	r := http.NewServeMux()
	r.HandleFunc("/", hiHandler)

	// Регистрация pprof-обработчиков
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
	// Вот и всё. Запустите приложение, а затем используйте pprof tool:
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

//// Чтобы начать настройку программы Go, нам необходимо включить профилирование.
//// Если бы код использовал поддержку бенчмаркинга пакета тестирования Go,
//// мы могли бы использовать стандартные флаги getest -cpuProfile и -memprofile.
//// В такой автономной программе, как эта, нам нужно импортировать runtime/pprof
//// и добавить несколько строк кода:
// var cpuProfile = flag.String("cpuProfile", "", "write cpu profile to file")

//func main() {
//fmt.Println("Hello, World!")
//flag.Parse()
//if *cpuProfile != "" {
//	f, err := os.Create(*cpuProfile)
//	if err != nil {
//		log.Fatal(err)
//	}
//	pprof.StartCPUProfile(f)
//	defer pprof.StopCPUProfile()
//}
//}
