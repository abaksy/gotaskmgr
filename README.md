# gotaskmgr
A simple task manager CLI in Go backed by BoltDB

# Usage

## Add a Task
Add a task to the database

```gotaskmgr add [taskname]```

example:

``` gotaskmgr add do the dishes```

Note: taskname can contain spaces

## List tasks
List tasks which are in-progress

```gotaskmgr list```

## List completed tasks
List tasks which are marked as completed

```gotaskmgr completed```

## Complete a task
Mark a task as completed in the database

```gotaskmgr do [taskID]```

Retreive the ID using the ```list``` command

## Remove a task
Remove a task from the database using its ID

```gotaskmgr rm [taskID]```