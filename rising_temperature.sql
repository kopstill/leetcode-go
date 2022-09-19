/*
    197: https://leetcode.cn/problems/rising-temperature/
 */

/* Write your MySQL query statement below */

/* datediff */
select a.id from weather a 
    cross join weather b 
    on datediff(a.recordDate, b.recordDate) = 1 
where a.temperature > b.temperature
