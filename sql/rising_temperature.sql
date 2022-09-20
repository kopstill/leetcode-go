/*
    197: https://leetcode.cn/problems/rising-temperature/
 */

/* Write your MySQL query statement below */

/* mysql: cross join == inner join == join */

/* datediff */
select a.id from weather a 
    cross join weather b 
    on datediff(a.recordDate, b.recordDate) = 1 
where a.temperature > b.temperature

/* timestampdiff */
select a.id from weather a 
    cross join weather b 
    on timestampdiff(day, a.recordDate, b.recordDate) = -1 
where a.temperature > b.temperature
