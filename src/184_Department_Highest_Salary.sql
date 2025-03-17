Create table If Not Exists Employee (id int, name varchar(255), salary int, departmentId int);
Create table If Not Exists Department (id int, name varchar(255));
    Truncate table Employee;
    insert into Employee (id, name, salary, departmentId) values ('1', 'Joe', '70000', '1');
    insert into Employee (id, name, salary, departmentId) values ('2', 'Jim', '90000', '1');
    insert into Employee (id, name, salary, departmentId) values ('3', 'Henry', '80000', '2');
    insert into Employee (id, name, salary, departmentId) values ('4', 'Sam', '60000', '2');
    insert into Employee (id, name, salary, departmentId) values ('5', 'Max', '90000', '1');
    Truncate table Department;
    insert into Department (id, name) values ('1', 'IT');
    insert into Department (id, name) values ('2', 'Sales');

with departmentInfo as
         (select Employee.departmentId, max(Employee.salary) as maxSalary, Department.name as Department
          from Department
                   left join Employee on Employee.departmentId = Department.id
          group by Employee.departmentId)

select departmentInfo.Department, Employee.name as Employee, Employee.salary as Salary
from Employee
         join departmentInfo
              on departmentInfo.departmentId = Employee.departmentId and departmentInfo.maxSalary = Employee.salary