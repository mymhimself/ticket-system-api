package postgres

const (
	QryTicketInfoFilter = `with tmp as(
							select distinct tt.*, 
							tm.sender_id, 
							min(case when tm.seen  then 1 else 0 end) over (PARTITION by thread_id) is_seen, 
							min(case when tm.replied then 1 else 0 end) over (PARTITION by thread_id) is_replied
							from ticket_threads tt, ticket_messages tm
							where tt.deleted_at is null and tt.id = tm.thread_id %s)
							select 
							id, 
							sender_id, 
							title, 
							priority, 
							department, 
							status, 
							bool(is_seen) seen,
							bool(is_replied) replied
							from tmp`

//where department = 1 and priority = 1 and status = 0 and is_seen = 0 and is_replied = 0

)
