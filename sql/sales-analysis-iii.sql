/*
    1084: https://leetcode.cn/problems/sales-analysis-iii/
 */

-- Write your MySQL query statement below

select product_id, product_name from (
    select p.product_id, p.product_name from Product p 
        where p.product_id in (
            select product_id from Sales 
                where sale_date between '2019-01-01' and '2019-03-31')
    ) t 
    where product_id not in (
        select product_id from Sales 
            where sale_date < '2019-01-01' or sale_date > '2019-03-31'
    );

select product_id, product_name from Product
    where product_id in (
        select product_id from Sales
            where sale_date between '2019-01-01' and '2019-03-31')
    and product_id not in (
        select product_id from Sales
            where sale_date < '2019-01-01' or sale_date > '2019-03-31'
    );

-- not between
select distinct a.product_id, b.product_name
    from sales a, product b
        where a.product_id = b.product_id
            and a.product_id not in (
                select product_id
                    from sales 
                        where sale_date 
                            not between '2019-01-01' and '2019-03-31'
            );

select distinct a.product_id, b.product_name
    from sales a inner join product b
        on a.product_id = b.product_id
            where a.product_id not in (
                select product_id
                    from sales 
                        where sale_date 
                            not between '2019-01-01' and '2019-03-31'
            );
