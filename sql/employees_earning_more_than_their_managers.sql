/*
    181: https://leetcode.cn/problems/employees-earning-more-than-their-managers/
 */

-- Write your MySQL query statement below

select a.name Employee from employee a join employee b on a.managerId = b.id where a.salary > b.salary;
select a.name Employee from employee a inner join employee b on a.managerId = b.id where a.salary > b.salary;
select a.name Employee from employee a cross join employee b on a.managerId = b.id where a.salary > b.salary;