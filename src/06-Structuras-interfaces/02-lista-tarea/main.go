package main

import "fmt"

// =============================================
// Lista de tareas
// =============================================
type TaskList struct {
	task []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.task = append(tl.task, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
}

// =============================================
// Tarea
// =============================================
type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

// =============================================

func main() {
	t1 := Task{
		name:      "Curso de Go Udemy",
		desc:      "Completar curso de Go durante este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de HTML youtube",
		desc:      "Completar curso de HTML durante este mes",
		completed: true,
	}

	t3 := Task{
		name:      "Curso de Css Coder",
		desc:      "Completar curso de Css durante este mes",
		completed: false,
	}

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)
	lista.appendTask(&t3)

	for i, task := range lista.task {
		fmt.Println(i, task.name)
	}

	lista.removeTask(0)
	for i, task := range lista.task {
		fmt.Println(i, task.name)
	}
}
