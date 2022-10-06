/*
    1667: https://leetcode.cn/problems/fix-names-in-a-table/
 */

-- Write your MySQL query statement below

-- substring
select 
    user_id, 
    CONCAT(Upper(left(name, 1)), Lower(substring(name, 2))) name
from 
    users
order by 
    user_id;

-- substr
select
    user_id,
    concat(upper(left(name, 1)), lower(substr(name, 2))) as name
from 
    Users
order by 
    user_id;

-- left, right
select 
    user_id, 
    CONCAT(UPPER(LEFT(name, 1)), LOWER(RIGHT(name, length(name) - 1))) as name
from 
    Users
order by 
    user_id;
