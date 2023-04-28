terraform {
  required_providers {
    todoist = {
      source = "hilyas.com/hilyas/todoist"
      version = "0.1.0"
    }
  }
}

provider "todoist" {
  api_key = "9e098c0362e2d5ee63576b64e0eeb3a94a21eeef"
}

##################################################
# Data Sources                                   #
##################################################

# data "todoist_project" "example_project" {
#   project_id = "2312101643"
# }

# data "todoist_task" "example_task" {
#   task_id = "6831856424"
# }

##################################################
# Resources                                      #
##################################################

resource "todoist_project" "example_project" {
  name        = "Example Project"
  parent_id   = ""
  color       = "red"
  is_favorite = true
  view_style  = "list"
}

resource "todoist_task" "example_task" {
  content      = "Example Task"
  # section_id   = ""
  description  = "This is an example task"
  is_completed = false
  labels       = ["urgent", "important"]
  priority     = 1
  project_id = todoist_project.example_project.id
  # due {
  #   string       = "2023-12-31"
  #   date         = "2023-12-31"
  #   is_recurring = false
  #   datetime     = "2023-09-01T12:00:00.000000Z"
  #   timezone     = "Europe/London"
  # }
}

##################################################
# Outputs                                        #
##################################################

# output "project_info" {
#   value = {
#     id          = data.todoist_project.example_project.id
#     name        = data.todoist_project.example_project.name
#     parent_id   = data.todoist_project.example_project.parent_id
#     color       = data.todoist_project.example_project.color
#     is_favorite = data.todoist_project.example_project.is_favorite
#     view_style  = data.todoist_project.example_project.view_style
#   }
# }

# output "task_info" {
#   value = {
#     id            = data.todoist_task.example_task.id
#     content       = data.todoist_task.example_task.content
#     project_id    = data.todoist_task.example_task.project_id
#     section_id    = data.todoist_task.example_task.section_id
#     description   = data.todoist_task.example_task.description
#     is_completed  = data.todoist_task.example_task.is_completed
#     labels        = data.todoist_task.example_task.labels
#     priority      = data.todoist_task.example_task.priority
#     # due           = data.todoist_task.example_task.due
#   }
# }