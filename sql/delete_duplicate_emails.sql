/*
    196: https://leetcode.cn/problems/delete-duplicate-emails/
 */

-- Please write a DELETE statement and DO NOT write a SELECT statement.
-- Write your MySQL query statement below

delete from Person where id not in (select id from (select min(id) id from Person p group by email) t);
