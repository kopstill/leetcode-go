/*
    1407: https://leetcode.cn/problems/top-travellers/
 */

-- Write your MySQL query statement below

-- ifnull
select u.name, ifnull(sum(r.distance), 0) as travelled_distance 
    from Rides r right join Users u 
        on r.user_id = u.id 
    group by r.user_id 
    order by travelled_distance desc, name;

-- if + isnull
select u.name, if(isnull(r.distance) = 1, 0, sum(r.distance)) as travelled_distance 
    from Rides r right join Users u 
        on r.user_id = u.id 
    group by r.user_id 
    order by travelled_distance desc, name asc;
