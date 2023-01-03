/*
    620: https://leetcode.cn/problems/not-boring-movies/
 */

-- Write your MySQL query statement below

select * from cinema where description <> "boring" and id % 2 = 1 order by rating desc;

select * from cinema where mod(id, 2) = 1 and description != 'boring' order by rating desc;

select * from cinema where id & 1 and description <> 'boring' order by rating desc;
