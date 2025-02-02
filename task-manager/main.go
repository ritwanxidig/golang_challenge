package main

import (
	"fmt"
	"task-manager/task"
)

	/* main is the entry point for the task manager application. It provides an interactive menu for the user to
	 * manage their tasks and mark them as done. The user can add up to 50 tasks. The application prompts the user
	 * to choose an option and based on the option, it either adds a task, views all tasks, marks a task, deletes a
	 * task, or exits the application. The application also keeps track of the total number of tasks the user has
	 * created. When the user reaches the maximum number of tasks (50), the application exits. */
func main() {
	// Initialize the current tasks count from the task package.
	currentTasksCount := len(task.Tasks)

	fmt.Println("Hello, Welcome to the Task Manager Application")
	fmt.Println("Here you can manage your tasks and Mark accordingly.")

	// Allow the user to add tasks until the total reaches 50.
	for currentTasksCount < 50 {
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add a new task")
		fmt.Println("2. View all tasks")
		fmt.Println("3. Mark a task")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading option:", err)
			continue
		}

		switch option {
		case 1:
			fmt.Println("Adding a new task...")
			fmt.Println("------------------")
			// Create a new task.
			description, done := getCreateUserInput()
			createdTask, err := task.Create(description, done)
			if err != nil {
				fmt.Println("Task creation failed:", err)
				continue
			}
			currentTasksCount++
			fmt.Println("Task created successfully:", createdTask)
			fmt.Printf("Total tasks: %d\n", currentTasksCount)
		case 2:
			fmt.Println("Viewing all tasks...")
			fmt.Println("------------------")
			// Retrieve and display all tasks.
			tasks, err := task.GetAll()
			if err != nil {
				fmt.Println("Failed to retrieve tasks:", err)
				continue
			}
			fmt.Println("Task List:")
			for _, t := range tasks {
				fmt.Printf("ID: %d, Description: %s, Done: %t\n", t.ID, t.Description, t.Done)
				fmt.Println("--------------------")
			}
			fmt.Printf("Total tasks: %d\n", currentTasksCount)
			fmt.Printf("--------------------------------END--------------------------------\n")
		case 3:
			fmt.Println("Marking a task...")
			fmt.Println("------------------")
			// Prompt the user for the task ID to Mark.
			taskID := getUserIdInput()
			if taskID == 0 {
				fmt.Println("Invalid task ID. Exiting Mark task process.")
				continue
			}
			var targetTask, err = task.GetByID(taskID)
			if err != nil {
				fmt.Println("Failed to retrieve task:", err)
				continue
			}

			fmt.Printf("Marking task: %s\n", targetTask.Description)
			fmt.Printf("Current status: %t\n", targetTask.Done)

			fmt.Printf("Do you ensure that you want to mark this task as done? (yes/no): ")

			var confirmationOption int = getConfirmInput() // 1 for yes "Done true", 2 for no "Done false"

			var _, error = task.Update(targetTask.Description, confirmationOption == 1, targetTask.ID)
			if error != nil {
				fmt.Println("Failed to update task:", error)
				continue
			}
			status := "Not Done"
			if confirmationOption == 1 {
				status = "Done"
			}
			fmt.Printf("Task marked as %s.\n", status)
			fmt.Printf("Total tasks: %d\n", currentTasksCount)
			fmt.Printf("--------------------------------END--------------------------------\n")
		case 4:
			fmt.Println("Deleting a task...")
			fmt.Println("------------------")
			// Prompt the user for the task ID to delete.
			taskID := getUserIdInput()
			if taskID == 0 {
				fmt.Println("Invalid task ID. Exiting Delete task process.")
				continue
			}
			var targetTask, err = task.GetByID(taskID)
			if err != nil {
				fmt.Println("Failed to retrieve task:", err)
				continue
			}

			fmt.Printf("Deleting task: %s\n", targetTask.Description)

			var confirmationOption int = getConfirmInput() // 1 for yes "Done true", 2 for no "Done false"
			if confirmationOption == 1 {
				var _, error = task.Delete(targetTask.ID)
				if error != nil {
					fmt.Println("Failed to delete task:", error)
					continue
				}
				currentTasksCount--
				fmt.Printf("Task deleted successfully.\n")
				fmt.Printf("Total tasks: %d\n", currentTasksCount)
				fmt.Printf("--------------------------------END--------------------------------\n")
			} else {
				fmt.Printf("Task deletion cancelled.\n")
			}
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

	fmt.Println("You have reached the maximum number of tasks (50). Exiting application.")
}
