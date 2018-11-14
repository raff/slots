# slots
A simple object to get empty slots in a list

My main use case for this was a number N of goroutines running T tasks (T > N) that would last different amount of time
(i.e. file uploads) and wanted to display each task progress on N lines of the screen (and re-assigning lines to tasks
when they become available)
