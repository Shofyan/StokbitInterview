select 
u.ID, 
u.UserName,
p.Username as ParentUserName 
from USER u 
left join USER p on u.Parent == p.ID
