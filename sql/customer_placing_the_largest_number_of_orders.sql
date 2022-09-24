/*
    586: https://leetcode.cn/problems/customer-placing-the-largest-number-of-orders/
 */

-- Write your MySQL query statement below

select customer_number from (
    select customer_number, count(*) as total from Orders 
        group by customer_number 
        order by total desc 
        limit 1
    ) t;
