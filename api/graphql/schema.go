package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

func NewSchema(dbInstance db.Querier) (graphql.Schema, error) {
	resolver := NewResolver(dbInstance)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetUserByID,
				},
				"users": &graphql.Field{
					Type:    graphql.NewList(UserType),
					Resolve: resolver.ResolveGetUsers,
				},
				"project": &graphql.Field{
					Type: graphql.NewList(ProjectType),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetProjectByID,
				},
				"projects": &graphql.Field{
					Type:    graphql.NewList(ProjectType),
					Resolve: resolver.ResolveGetProjects,
				},
				"assignedProject": &graphql.Field{
					Type: AssignedProjectType,
					Args: graphql.FieldConfigArgument{
						"project_id": &graphql.ArgumentConfig{Type: graphql.Int},
					},
					Resolve: resolver.ResolveGetAssignedProjectByID,
				},
				"assignedProjects": &graphql.Field{
					Type:    graphql.NewList(AssignedProjectType),
					Resolve: resolver.ResolveGetAssignedProjects,
				},
				"ProjectCommentsByProjectID": &graphql.Field{
					Type: graphql.NewList(ProjectCommentType),
					Args: graphql.FieldConfigArgument{
						"project_id": &graphql.ArgumentConfig{Type: graphql.Int},
					},
					Resolve: resolver.ResolveGetProjectCommentsByProjectID,
				},
				"ProjectCommentsByID": &graphql.Field{
					Type: ProjectCommentType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.Int},
					},
					Resolve: resolver.ResolveGetProjectCommentByID,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				"insertUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertUserInputType)},
					},
					Resolve: resolver.ResolveInsertUser,
				},
				"deleteUser": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteUser,
				},
				"updateUser": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateUserInputType)},
					},
					Resolve: resolver.ResolveUpdateUser,
				},
				"insertProject": &graphql.Field{
					Type: ProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertProjectInputType)},
					},
					Resolve: resolver.ResolveInsertProject,
				},
				"deleteProject": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteProject,
				},
				"updateProject": &graphql.Field{
					Type: ProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateProjectInputType)},
					},
					Resolve: resolver.ResolveUpdateProject,
				},
				"insertAssignedProject": &graphql.Field{
					Type: AssignedProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertAssignedProjectInputType)},
					},
					Resolve: resolver.ResolveInsertAssignedProject,
				},
				"deleteAssignedProject": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"project_id": &graphql.ArgumentConfig{Type: graphql.Int},
					},
					Resolve: resolver.ResolveDeleteAssignedProject,
				},
				"updateAssignedProject": &graphql.Field{
					Type: AssignedProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateAssignedProjectInputType)},
					},
					Resolve: resolver.ResolveUpdateAssignedProject,
				},
				"insertProjectComment": &graphql.Field{
					Type: ProjectCommentType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertProjectCommentInputType)},
					},
					Resolve: resolver.ResolveInsertProjectComment,
				},
				"deleteProjectComment": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteProject,
				},
				"updateProjectComment": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateProjectCommentInputType)},
					},
					Resolve: resolver.ResolveUpdateProjectComment,
				},
			},
		}),
	})
}
