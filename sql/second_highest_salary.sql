/*
    176: https://leetcode.cn/problems/second-highest-salary/
 */

/* Write your MySQL query statement below */

select (select distinct salary from employee order by salary desc limit 1 offset 1) as SecondHighestSalary

select ifnull((select distinct salary from employee order by salary desc limit 1 offset 1), null) as SecondHighestSalary