package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}

	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de Platzi en esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de NodeJS",
		descripcion: "Completar mi curso de NodeJS de Platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarALista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	//fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Nestor"] = lista

	t4 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi curso de Java de Platzi en esta semana",
	}

	t5 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# de Platzi en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Ricardo"] = lista2

	fmt.Println("Tareas de Nestor")
	mapaTareas["Nestor"].imprimirLista()
	fmt.Println("Tareas de Ricardo")
	mapaTareas["Ricardo"].imprimirLista()

}
