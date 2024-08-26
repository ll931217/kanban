package main

import "github.com/charmbracelet/bubbles/list"

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}

	b.cols[todo].list.Title = "To Do"
	b.cols[todo].list.SetItems([]list.Item{
		Task{title: "Task 1", description: "Do something else", status: todo},
	})

	b.cols[inProgress].list.Title = "In Progress"
	b.cols[inProgress].list.SetItems([]list.Item{
		Task{title: "Task 2", description: "Do something else", status: inProgress},
	})

	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		Task{title: "Task 3", description: "Do something else again", status: done},
	})
}
