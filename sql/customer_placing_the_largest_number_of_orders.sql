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

SELECT
    customer_number
FROM
    orders
GROUP BY customer_number
ORDER BY COUNT(*) DESC
LIMIT 1
;

-- compatible with multiple identical records
select customer_number from orders
group by customer_number
having count(*) = (
    select count(*) from orders
    group by customer_number
    order by count(*) desc
    limit 1
);
