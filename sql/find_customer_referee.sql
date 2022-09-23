/*
    584: https://leetcode.cn/problems/find-customer-referee/
 */

-- Write your MySQL query statement below

select name from customer where referee_id != 2 or referee_id is null;
select name from customer where referee_id <> 2 or referee_id is null;
select name from customer where id not in (select id from customer where referee_id = 2);
select name from customer where ifnull(referee_id, -1) != 2;
SELECT name FROM customer where not referee_Id <=> 2;
