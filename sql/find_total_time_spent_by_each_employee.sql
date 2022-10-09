/*
    1741: https://leetcode.cn/problems/find-total-time-spent-by-each-employee/
 */

-- Write your MySQL query statement below

select 
    event_day day, 
    emp_id, 
    sum(out_time - in_time) total_time 
from 
    Employees 
group by 
    emp_id, 
    event_day;
