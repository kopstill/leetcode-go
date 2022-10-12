/*
    1890: https://leetcode.cn/problems/the-latest-login-in-2020/
 */

-- Write your MySQL query statement below

select 
    user_id, 
    max(time_stamp) as last_stamp 
from 
    Logins
where 
    time_stamp between '2020-01-01 00:00:00' and '2020-12-31 23:59:59'
group by user_id

select 
    user_id, 
    max(time_stamp) as last_stamp 
from Logins
where year(time_stamp) = 2020
group by user_id

select 
    user_id, max(time_stamp) as last_stamp 
from Logins
where time_stamp between '2020-01-01' and '2021-01-01'
group by user_id

select 
    user_id, max(time_stamp) as last_stamp 
from Logins
where time_stamp rlike '^2020'
group by user_id
