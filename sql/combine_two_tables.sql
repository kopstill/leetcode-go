/*
    175: https://leetcode.cn/problems/combine-two-tables/
 */

/* Write your MySQL query statement below */

select p.firstName, p.lastName, a.city, a.state from Person p left join Address a on p.personId = a.personId;

select 
    FirstName, 
    LastName, 
    (select City from Address a where a.PersonId = p.PersonId) as City, 
    (select State from Address a where a.PersonId = p.PersonId) as State 
from Person p;
