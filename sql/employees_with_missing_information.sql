/*
    1965: https://leetcode.cn/problems/employees-with-missing-information/
 */

-- Write your MySQL query statement below

select employee_id from (
    select e.employee_id, e.name, s.salary from Employees e left join Salaries s on e.employee_id = s.employee_id
    union all
    select s.employee_id, e.name, s.salary from Employees e right join Salaries s on e.employee_id = s.employee_id
) t 
where 
    t.name is null 
    or 
    t.salary is null 
order by employee_id;

select e.employee_id employee_id from Employees e 
    left join Salaries s on e.employee_id = s.employee_id where salary is null
union all
select s.employee_id employee_id from Employees e 
    right join Salaries s on e.employee_id = s.employee_id where name is null
order by employee_id;

SELECT
	t.employee_id 
FROM
	( SELECT employee_id FROM employees UNION SELECT employee_id FROM salaries ) t
	LEFT JOIN employees e ON e.employee_id = t.employee_id
	LEFT JOIN salaries s ON s.employee_id = t.employee_id 
WHERE
	e.NAME IS NULL 
	OR s.salary IS NULL 
ORDER BY
	t.employee_id;

select 
    employee_id 
from 
(
    select employee_id from employees
    union all 
    select employee_id from salaries
) as t
group by 
    employee_id
having 
    count(employee_id) = 1
order by 
    employee_id;

select e.employee_id employee_id from Employees e left join Salaries s on e.employee_id=s.employee_id where s.employee_id is null
    union all
select s.employee_id employee_id from Employees e right join Salaries s on e.employee_id=s.employee_id where e.employee_id is null
order by employee_id;

select e.employee_id employee_id from Employees e where not exists (select 1 from Salaries s where s.employee_id = e.employee_id)
union all
select s.employee_id employee_id from Salaries s where not exists (select 1 from Employees e where s.employee_id = e.employee_id)
order by employee_id;

select employee_id from Employees where employee_id not in (select employee_id from Salaries)
union all
select employee_id from Salaries where employee_id not in (select employee_id from Employees)
order by employee_id;
