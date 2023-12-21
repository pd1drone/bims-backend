CREATE DATABASE bims;

USE bims;

CREATE TABLE `Users` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `FullName` varchar(255),
  `FirstName` varchar(255),
  `MiddleName` varchar(255),
  `LastName` varchar(255),
  `PositionID` integer,
  `Email` varchar(255),
  `Username` varchar(255),
  `Password` varchar(255),
  `IsAdmin` bool,
  `ProfileLink` varchar(255)
);

CREATE TABLE `Positions` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Name` varchar(255),
  `Description` varchar(255)
);

CREATE TABLE `Residents` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `LastName` varchar(255),
  `FirstName` varchar(255),
  `MiddleName` varchar(255),
  `Address` varchar(255),
  `BirthDate` varchar(255),
  `BirthPlace` varchar(255),
  `Gender` varchar(255),
  `CivilStatus` varchar(255),
  `ContactNumber` varchar(255),
  `GuardianName` varchar(255),
  `GurdianContactNumber` varchar(255),
  `IssuingOfficer` varchar(255),
  `DocumentType` varchar(255),
  `DocumentID` int
);

CREATE TABLE `Clearance` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `ResidentID` integer,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `ValidUntil` varchar(255),
  `IssuingOfficer` varchar(255),
  `Remarks` longtext,
  `ResidentLastName` varchar(255),
  `ResidentFirstName` varchar(255),
  `ResidentMiddleName` varchar(255),
  `Purpose` longtext,
  `CedulaNo` varchar(255),
  `PrecintNo` varchar(255),
  `DocumentStatus` varchar(255)
);

CREATE TABLE `Referrals` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `ResidentID` integer,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `HCGGGNumber` varchar(255),
  `PhilHealthID` varchar(255),
  `PhilHealthCategory` varchar(255),
  `ReasonForReferral` longtext,
  `ValidUntil` varchar(255),
  `IssuingOfficer` varchar(255),
  `Remarks` longtext,
  `DocumentStatus` varchar(255)
);

CREATE TABLE `Indigencies` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `ResidentID` integer,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `Reason` longtext,
  `ValidUntil` varchar(255),
  `IssuingOfficer` varchar(255),
  `Remarks` longtext,
  `DocumentStatus` varchar(255)
);

CREATE TABLE `BDRRMC` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `TypeOfRecord` varchar(255),
  `PartiesInvolved` varchar(255),
  `DateTime` varchar(255),
  `Location` varchar(255),
  `RecordDetails` longtext,
  `IssuingOfficer` varchar(255)
);

CREATE TABLE `Incidents` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `DateCreated` varchar(255),
  `DateUpdated` varchar(255),
  `CompliantFullName` varchar(255),
  `Respondent` varchar(255),
  `IncidentStatus` varchar(255),
  `IncidentDateTime` varchar(255),
  `IncidentLocation` varchar(255),
  `IncidentNarration` longtext,
  `IssuingOfficer` varchar(255)
);


-- insert barangay positions
INSERT INTO Positions (Name,Description)VALUES('admin','admin');
INSERT INTO Positions (Name,Description)VALUES('punong_barangay','Punong Barangay');
INSERT INTO Positions (Name,Description)VALUES('kagawad','Kagawad');
INSERT INTO Positions (Name,Description)VALUES('midwife','Midwife');
INSERT INTO Positions (Name,Description)VALUES('bhw','BHW');
INSERT INTO Positions (Name,Description)VALUES('ex_o','Ex-O');
INSERT INTO Positions (Name,Description)VALUES('staff','Barangay Staff');

-- username= admin , password = admin123
INSERT INTO Users (FullName,FirstName,MiddleName,LastName,PositionID,Email,Username, Password, IsAdmin, ProfileLink) VALUES ('admin','admin','admin','admin',1,'admin@admin.com','admin','0192023a7bbd73250516f069df18b500',true,'');