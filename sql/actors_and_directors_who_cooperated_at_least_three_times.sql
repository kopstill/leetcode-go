/*
    1050: https://leetcode.cn/problems/actors-and-directors-who-cooperated-at-least-three-times/
 */

-- Write your MySQL query statement below

select actor_id, director_id from ActorDirector group by actor_id, director_id having count(*) >= 3;
