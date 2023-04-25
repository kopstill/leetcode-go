/*
    1789: https://leetcode.cn/problems/primary-department-for-each-employee/
 */

-- Write your MySQL query statement below
select * from (
    select employee_id, department_id from Employee where employee_id in (select employee_id from Employee group by employee_id having count(*) > 1) and primary_flag = 'Y'
    union all
    select employee_id, department_id from Employee where employee_id in (select employee_id from Employee group by employee_id having count(*) = 1)
) t order by t.employee_id
