/*
    182: https://leetcode.cn/problems/duplicate-emails/
 */

-- Write your MySQL query statement below

select Email from Person group by Email having count(*) > 1;
select Email from Person group by Email having count(Email) > 1;
