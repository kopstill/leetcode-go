/*
  619: https://leetcode.cn/problems/biggest-single-number/
 */

-- Write your MySQL query statement below

-- subquery
select (select num from MyNumbers group by num having count(*) = 1 order by num desc limit 1) as num;
select (select num from MyNumbers group by num having count(*) = 1 order by num desc limit 1) as num from dual;

-- max function
select max(num) as num from (select num from MyNumbers group by num having count(num) = 1) as t;
