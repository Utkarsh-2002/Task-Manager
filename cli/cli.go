package cli

import (
	"fmt"
	"os"
	"strconv"
	"TASK-MANAGER/handlers"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: "task-manager"}      // The root command is the top-level command for the application. in this case, it's task-manager this is the command that users run to invoke the CLI tool.

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new task",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide a title and description for the task.")
				os.Exit(1)
			}
			title := args[0]
			description := args[1]

			// Create task through CLI function
			task, err := handlers.CreateTaskCLI(title, description)
			if err != nil {
				fmt.Println("Error creating task:", err)
				return
			}
			fmt.Printf("Task created: %+v\n", task)
		},
	}

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get task by ID",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide a task ID.")
				os.Exit(1)
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID:", args[0])
				return
			}

			// Get task by ID through CLI function
			task, err := handlers.GetTaskByIDCLI(id)
			if err != nil {
				fmt.Println("Error getting task:", err)
				return
			}
			fmt.Printf("Task retrieved: %+v\n", task)
		},
	}

	rootCmd.AddCommand(createCmd, getCmd)     // This adds the createcmd and getcmd as subcommands to the root command(task-manager)
	rootCmd.Execute()
}
