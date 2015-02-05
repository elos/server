package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/models/event"
	"github.com/elos/models/task"
	"github.com/elos/models/user"
)

const (
	DataVersion           = 1
	UserKind    data.Kind = "user"
	EventKind   data.Kind = "event"
	TaskKind    data.Kind = "task"
)

const (
	UserEvents       data.LinkName = "events"
	UserTasks        data.LinkName = "tasks"
	UserCurrentTask  data.LinkName = "current_task"
	EventUser        data.LinkName = "user"
	TaskUser         data.LinkName = "user"
	TaskDependencies data.LinkName = "dependencies"
)

var RMap data.RelationshipMap = data.RelationshipMap{
	UserKind: {
		UserEvents: data.Link{
			Name:    UserEvents,
			Kind:    data.MulLink,
			Other:   EventKind,
			Inverse: EventUser,
		},
		UserTasks: data.Link{
			Name:    UserTasks,
			Kind:    data.MulLink,
			Other:   TaskKind,
			Inverse: TaskUser,
		},
		UserCurrentTask: data.Link{
			Name:  UserCurrentTask,
			Kind:  data.OneLink,
			Other: TaskKind,
		},
	},
	EventKind: {
		EventUser: data.Link{
			Name:    EventUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserEvents,
		},
	},
	TaskKind: {
		TaskUser: data.Link{
			Name:    TaskUser,
			Kind:    data.OneLink,
			Other:   UserKind,
			Inverse: UserTasks,
		},
		TaskDependencies: data.Link{
			Name:  TaskDependencies,
			Kind:  data.MulLink,
			Other: TaskKind,
		},
	},
}

func SetupStore(addr string) data.Store {
	mongo.RegisterKind(UserKind, "users")
	mongo.RegisterKind(EventKind, "events")
	mongo.RegisterKind(TaskKind, "tasks")

	db, err := mongo.NewDB(addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Database connection established")

	sch, err := data.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Schema successfully validated")

	s := data.NewStore(db, sch)

	s.Register(UserKind, user.NewM)
	s.Register(EventKind, event.NewM)
	s.Register(TaskKind, task.NewM)

	user.Setup(sch, UserKind, 1)
	user.Events = UserEvents
	user.Tasks = UserTasks
	user.CurrentTask = UserCurrentTask

	event.Setup(sch, EventKind, 1)
	event.User = EventUser

	task.Setup(sch, TaskKind, 1)
	task.User = TaskUser
	task.Dependencies = TaskDependencies

	return s
}
