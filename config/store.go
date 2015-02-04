package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/models/event"
	"github.com/elos/server/models/task"
	"github.com/elos/server/models/user"
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

func (s *Server) SetupStore(addr string) {
	mongo.RegisterKind(UserKind, "users")
	mongo.RegisterKind(EventKind, "events")
	mongo.RegisterKind(TaskKind, "tasks")

	db, err := mongo.NewDB(addr)
	if err != nil {
		log.Fatal(err)
	}

	Log("Database connection established")

	sch, err := data.NewSchema(&RMap, DataVersion)
	if err != nil {
		log.Fatal(err)
	}

	Log("Schema successfully validated")

	s.Store = data.NewStore(db, sch)

	s.Store.Register(UserKind, user.NewM)
	s.Store.Register(EventKind, event.NewM)
	s.Store.Register(TaskKind, task.NewM)

	user.Setup(sch, UserKind, 1)
	user.Events = UserEvents
	user.Tasks = UserTasks
	user.CurrentTask = UserCurrentTask

	event.Setup(sch, EventKind, 1)
	event.User = EventUser

	task.Setup(sch, TaskKind, 1)
	task.User = TaskUser
	task.Dependencies = TaskDependencies
}
