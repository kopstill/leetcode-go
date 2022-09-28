/*
    1141: https://leetcode.cn/problems/user-activity-for-the-past-30-days-i/
 */

-- Write your MySQL query statement below

select activity_date as day, count(distinct user_id) as active_users from Activity 
    where activity_date 
        between date_sub('2019-07-27', interval 29 day) and '2019-07-27' 
    group by activity_date;

-- datediff
select activity_date day, count(distinct user_id) active_users from activity
    where datediff('2019-07-27', activity_date) >= 0 
        and datediff('2019-07-27', activity_date) < 30
    group by activity_date;
