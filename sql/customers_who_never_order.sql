/*
    183: https://leetcode.cn/problems/customers-who-never-order/
 */

-- Write your MySQL query statement below

-- not in
select name Customers from Customers where Id not in (select CustomerId from Orders);
select c.Name Customers from Customers c where c.Id not in (select o.CustomerId from Orders o);

select customers.name as 'Customers' from customers
where customers.id not in
(
    select customerid from orders
);

-- not exists
select t1.Name AS Customers from Customers AS t1 
where not exists (select 1 from Orders AS t2 where t1.id = t2.CustomerId);

-- left join
select c.Name as Customers from Customers as c
    left join Orders as o on c.Id = o.CustomerId
where o.Id is null;
