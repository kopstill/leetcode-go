/*
    183: https://leetcode.cn/problems/customers-who-never-order/
 */

-- Write your MySQL query statement below

select c.Name Customers from Customers c where c.Id not in (select o.CustomerId from Orders o);

select customers.name as 'Customers' from customers
where customers.id not in
(
    select customerid from orders
);
