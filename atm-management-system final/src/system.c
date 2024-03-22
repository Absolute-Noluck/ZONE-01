#include "header.h"

const char *RECORDS = "./data/records.txt";

int getAccountFromFile(FILE *ptr, char name[50], struct Record *r)
{
    return fscanf(ptr, "%d %d %s %d %d/%d/%d %s %d %lf %s",
                  &r->id,
                  &r->userId,
                  name,
                  &r->accountNbr,
                  &r->deposit.month,
                  &r->deposit.day,
                  &r->deposit.year,
                  r->country,
                  &r->phone,
                  &r->amount,
                  r->accountType) != EOF;
}

void saveAccountToFile(FILE *ptr, struct User *u, struct Record *r)
{
    fprintf(ptr, "%d %d %s %d %d/%d/%d %s %d %.2lf %s\n\n",
            r->id,
            r->userId,
            u->name,
            r->accountNbr,
            r->deposit.month,
            r->deposit.day,
            r->deposit.year,
            r->country,
            r->phone,
            r->amount,
            r->accountType);
}

void stayOrReturn(int notGood, struct User u)
{
    int option;
    if (notGood == 0)
    {
        system("clear");
        printf("\n✖ Record not found!!\n");
    invalid:
        printf("\nEnter 0 to try again, 1 to return to main menu and 2 to exit:");
        scanf("%d", &option);
        if (option == 0)
        
            mainMenu(u);
        else if (option == 1)
            mainMenu(u);
        else if (option == 2)
            exit(0);
        else
        {
            printf("Insert a valid operation!\n");
            goto invalid;
        }
    }
    else
    {
        printf("\nEnter 1 to go to the main menu and 0 to exit:");
        scanf("%d", &option);
    }
    if (option == 1)
    {
        system("clear");
        mainMenu(u);
    }
    else
    {
        system("clear");
        exit(1);
    }
}

void success(struct User u)
{
    int option;
    printf("\n✔ Success!\n\n");
invalid:
    printf("Enter 1 to go to the main menu and 0 to exit!\n");
    scanf("%d", &option);
    system("clear");
    if (option == 1)
    {
        mainMenu(u);
    }
    else if (option == 0)
    {
        exit(1);
    }
    else
    {
        printf("Insert a valid operation!\n");
        goto invalid;
    }
}

void createNewAcc(struct User u)
{
    struct Record r;
    int temp = 0;
    struct Record cr;
    char userName[50];
    FILE *pf = fopen(RECORDS, "a+");

noAccount:
    system("clear");
    printf("\t\t\t===== New record =====\n");

    printf("\nEnter today's date(mm/dd/yyyy):");
    scanf("%d/%d/%d", &r.deposit.month, &r.deposit.day, &r.deposit.year);
    printf("\nEnter the account number:");
    scanf("%d", &r.accountNbr);

    while (getAccountFromFile(pf, userName, &cr))
    {
        if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
        {
            printf("✖ This Account already exists for this user\n\n");
            goto noAccount;
        }
        if (temp < cr.id)
        {
            temp = cr.id;
        }
    }
    r.id = temp + 1;
    r.userId = GetID(u);
    printf("\nEnter the country:");
    scanf("%s", r.country);
    printf("\nEnter the phone number:");
    scanf("%d", &r.phone);
    printf("\nEnter amount to deposit: $");
    scanf("%lf", &r.amount);
    printf("\nChoose the type of account:\n\t-> saving\n\t-> current\n\t-> fixed01(for 1 year)\n\t-> fixed02(for 2 years)\n\t-> fixed03(for 3 years)\n\n\tEnter your choice:");
    scanf("%s", r.accountType);

    saveAccountToFile(pf, &u, &r);

    fclose(pf);
    success(u);
}

void checkAllAccounts(struct User u)
{
    char userName[100];
    struct Record r;

    FILE *pf = fopen(RECORDS, "r");

    system("clear");
    printf("\t\t====== All accounts from user, %s =====\n\n", u.name);
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0)
        {
            printf("_____________________\n");
            printf("\nAccount number:%d\nDeposit Date:%d/%d/%d \ncountry:%s \nPhone number:%d \nAmount deposited: $%.2f \nType Of Account:%s\n",
                   r.accountNbr,
                   r.deposit.day,
                   r.deposit.month,
                   r.deposit.year,
                   r.country,
                   r.phone,
                   r.amount,
                   r.accountType);
        }
    }
    fclose(pf);
    success(u);
}
void RemoveAccount(struct User u, int account_number)
{
    char userName[100];
    struct Record r;
    int i;
    char filename[40] = "./data/records.txt";
    FILE *fileptr1, *fileptr2;
    if ((fileptr1 = fopen(filename, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    int check;
    while (getAccountFromFile(fileptr1, userName, &r))
    {

        if (strcmp(userName, u.name) == 0)
        {
            if (r.accountNbr == account_number)
            {
                printf("_____________________\n");
                check = 1;
                break;
            }
        }
    }
    if (check == 1)
    {

        struct User userChecker;

        char ch;

        int temp = 0;

        fileptr2 = fopen("./data/replica.c", "w");
        rewind(fileptr1);
        while (getAccountFromFile(fileptr1, userName, &r))
        {
            if (strcmp(userName, u.name) == 0 && (r.accountNbr == account_number))
            {
                continue;
            }
            else
            {
                userChecker.id = r.userId;
                memcpy(userChecker.name, userName, sizeof(userName));
                saveAccountToFile(fileptr2, &userChecker, &r);
            }
        }
        fclose(fileptr1);
        fclose(fileptr2);
        rename("./data/replica.c", filename);
    }
    else
    {
        printf("Account does not exist\n");
        return;
    }
}

void RemoveExistingUser(struct User u)
{
    FILE *fileptr1, *fileptr2;
    struct User userChecker;
    char filename[40] = "./data/users.txt";
    char ch;
    int count = 0;
    int temp = 0;

    if ((fileptr1 = fopen(filename, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    rewind(fileptr1);
    while (fscanf(fileptr1, "%d %s %s", &userChecker.id, userChecker.name, userChecker.password) != EOF) 
    {

        if (strcmp(userChecker.password, u.password) == 0)
        {
            break;
        }
        count++;
    }
    fileptr2 = fopen("./data/replica.c", "w");
    rewind(fileptr1);
    while (ch != EOF)
    {
        ch = getc(fileptr1);
        if (ch == '\n')
        {
            temp++;
        }
        if (temp != count)
        {
            if (ch != EOF)
            {
                putc(ch, fileptr2);
            }
        }

    }
    fclose(fileptr1);
    fclose(fileptr2);
    remove(filename);
    rename("./data/replica.c", filename); 
    success(u);
}
void CheckingDetailAccount(struct User u)
{
    char userName[100];
    struct Record r;
    int account_number;
    FILE *pf = fopen(RECORDS, "r");
    printf("\nAccount Numbers : \n");
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0)
        {

            printf("%d\n",
                   r.accountNbr);
        }
    }

    printf("\nEnter the Account Number:");
    scanf("%d", &account_number);
    if (CheckIfAccountIsOwnedByUser(u, u.name, account_number) == 0)
    {
        printf("Wrong Account Number");
        stayOrReturn(1, u);
        return;
    }
    rewind(pf);
    int count;
    int check;
    double result;
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0)
        {

            if (r.accountNbr == account_number)
            {
                printf("\nType Of Account:%s\n",
                       r.accountType);

                if (strcmp(r.accountType, "saving") == 0)
                {
                    printf("\ninterest rate 0.07%%\n");
                    result = (r.amount * 0.07) / 12;
                    printf("\ninterest:$%.2f every %d\n", result, r.deposit.day);
                }
                else if (strcmp(r.accountType, "current") == 0)
                {
                    printf("You will not get interests because the account is of type current");
                }
                else if (strcmp(r.accountType, "fixed01") == 0)
                {
                    printf("\ninterest rate 0.04%%\n");
                    result = (r.amount * 0.04);
                    printf("\ninterest:$%.2f on %d/%d/%d\n", result, r.deposit.day, r.deposit.month, r.deposit.year + 1);
                }
                else if (strcmp(r.accountType, "fixed02") == 0)
                {
                    printf("\ninterest rate 0.05%%\n");
                    result = (r.amount * 0.05) * 2;
                    printf("\ninterest:$%.2f on %d/%d/%d\n", result, r.deposit.day, r.deposit.month, r.deposit.year + 2);
                }
                else if (strcmp(r.accountType, "fixed03") == 0)
                {
                    printf("\ninterest rate 0.08%%\n");
                    result = (r.amount * 0.08) * 3;
                    printf("\ninterest:$%.2f on %d/%d/%d\n", result, r.deposit.day, r.deposit.month, r.deposit.year + 3);
                }
                printf("_____________________\n");

                break;
            }
        }
    }
    success(u);
}

void ModifyAccount(struct User u)
{
    char userName[100];
    struct Record r;
    struct Record cr;
    FILE *pf = fopen(RECORDS, "r");

    system("clear");
    printf("\t\t\t===== New record =====\n");
    printf("\nEnter the account number:");
    scanf("%d", &r.accountNbr);
    if (CheckIfAccountIsOwnedByUser(u, u.name, r.accountNbr) == 0)
    {
        printf("\nAccount does not exist");
        stayOrReturn(1, u);
        return;
    }

    while (getAccountFromFile(pf, userName, &cr))
    {
        if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
        {
            printf("\nEnter the country:");
            scanf("%s", cr.country);
            printf("\nEnter the phone number:");
            scanf("%d", &cr.phone);
            RemoveAccount(u, r.accountNbr);

            break;
        }
    }

    fclose(pf);
    FILE *write = fopen(RECORDS, "a+");
    saveAccountToFile(write, &u, &cr);
    fclose(write);
    success(u);
}
void TransferOwner(struct User u)
{
    char userName[100];
    struct Record r;
    struct Record cr;
    struct User TempUser;
    struct User Temp2User;
    FILE *pff = fopen(RECORDS, "r+");

    system("clear");
    printf("\t\t\t===== New record =====\n");

    printf("\nEnter the account name for the next owner:");
    scanf("%s", TempUser.name);

    u.id = GetID(TempUser);
    memcpy(Temp2User.name, u.name, 32);
    printf("\nEnter your account number to transfer:\n");
    scanf("%d", &r.accountNbr);
    if (CheckIfAccountIsOwnedByUser(u, u.name, r.accountNbr) == 0)
    {
        printf("\nAccount does not exist");
        stayOrReturn(1, u);
        return;
    }
    FILE *writee = fopen(RECORDS, "a+");

    while (getAccountFromFile(pff, userName, &cr))
    {
        if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
        {

            cr.userId = u.id;
            memcpy(u.name, TempUser.name, 32);
            saveAccountToFile(writee, &u, &cr);
            fclose(writee);                    
            memcpy(u.name, Temp2User.name, 32);
            RemoveAccount(u, r.accountNbr);
        }
    }


    fclose(pff);
    success(u);
}
int CheckIfUserExist(struct User u, char userName[50])
{
    struct Record r;
    FILE *pf = fopen(RECORDS, "r");             
    while (getAccountFromFile(pf, userName, &r)) 
    {
        if (strcmp(userName, u.name) == 0)
        {
            fclose(pf); 
            return 1;
        }
    }
    fclose(pf);
    return 0;
}

int CheckIfAccountIsOwnedByUser(struct User u, char userName[50], int account_number)
{
    struct Record r;
    FILE *pf = fopen(RECORDS, "r");             
    while (getAccountFromFile(pf, userName, &r))
    {
        if (strcmp(userName, u.name) == 0 && r.accountNbr == account_number)
        {
            fclose(pf);
            return 1;
        }
    }
    fclose(pf);
    return 0;
}

int GetID(struct User u)
{
    FILE *fileptr1, *fileptr2;
    struct User userChecker;
    char filename[40] = "./data/users.txt";
    char ch;
    int count = 0;
    int temp = 0;

    if ((fileptr1 = fopen(filename, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    rewind(fileptr1);
    while (fscanf(fileptr1, "%d %s %s", &userChecker.id, userChecker.name, userChecker.password) != EOF)
    {
        if (strcmp(userChecker.name, u.name) == 0)
        {
            return userChecker.id;
        }
    }
}
int AutoIncrementID()
{
    FILE *fileptr1;
    struct User userChecker;
    char filename[40] = "./data/users.txt";
    int count = 0;

    if ((fileptr1 = fopen(filename, "r")) == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    rewind(fileptr1);
    while (fscanf(fileptr1, "%d %s %s", &userChecker.id, userChecker.name, userChecker.password) != EOF)
    {
        if (userChecker.id > count)
        {
            count = userChecker.id;
        }
    }
    return count + 1;
}

void MakeTransaction(struct User u)
{
    char userName[100];
    char TypeAccont[100];
    struct Record r;
    struct Record cr;
    struct User TempUser;
    double account_amount;
    FILE *pff = fopen(RECORDS, "r");
    memcpy(TempUser.name, u.name, 32);
    int transactionType;
    system("clear");
    printf("\t\t\t===== New record =====\n");
    printf("0: Deposit\n1: Withdraw\n2: Transfer\n3: Back\n");
    scanf("%d", &transactionType);
    if (transactionType == 0)
    {
        printf("\nEnter the account number to deposit:");
        scanf("%d", &r.accountNbr);
        printf("\nEnter the amount to deposit:");
        scanf("%lf", &account_amount);

        if (CheckIfAccountIsOwnedByUser(u, u.name, r.accountNbr) == 0)
        {
            printf("\nAccount does not exist");
            stayOrReturn(1, u);
            return;
        }
        while (getAccountFromFile(pff, userName, &cr))
        {

            if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
            {
                fclose(pff);
                break;
            }
        }
        if (strcmp(cr.accountType, "fixed01") == 0 || strcmp(cr.accountType, "fixed02") == 0 || strcmp(cr.accountType, "fixed03") == 0)
        { //Verify if the AccountType is "fixed01", "fixed02" or "fixed03"
            printf("\nError\n");
            return;
        }
        RemoveAccount(u, cr.accountNbr);
        cr.amount += account_amount;

        printf("%lf", cr.amount);
        printf("%d", cr.accountNbr);

        FILE *writee = fopen(RECORDS, "a+");
        saveAccountToFile(writee, &u, &cr);
        fclose(writee);
        success(u);
    }
    else if (transactionType == 1)
    {
        printf("\nEnter the account number to withdraw:");
        scanf("%d", &r.accountNbr);
        printf("\nEnter the amount to withdraw:");
        scanf("%lf", &account_amount);
        if (CheckIfAccountIsOwnedByUser(u, u.name, r.accountNbr) == 0)
        {
            printf("\nAccount does not exist");
            stayOrReturn(1, u);
            return;
        }
        while (getAccountFromFile(pff, userName, &cr))
        {
            if (strcmp(userName, u.name) == 0 && cr.accountNbr == r.accountNbr)
            {
                fclose(pff);
                break;
            }
        }
        if (cr.amount < account_amount)
        {
            printf("\nInsufficient funds");
            stayOrReturn(1, u);
            return;
        }
        cr.amount -= account_amount;
        RemoveAccount(u, r.accountNbr);
        FILE *writee = fopen(RECORDS, "a+");
        saveAccountToFile(writee, &u, &cr);
        fclose(writee);
        success(u);
    }
    else
    {
        while (getAccountFromFile(pff, userName, &cr))
        {
            if (strcmp(userName, u.name) == 0)
            {

                printf("\nEnter the account number:");
                scanf("%d", &r.accountNbr);
                printf("\nEnter the amount to transfer:");
                scanf("%lf", &account_amount);
                printf("\nEnter the account number to transfer to:");
                scanf("%d", &cr.accountNbr);
                GetAccountByUser(cr.accountNbr, &cr, u);
                GetAccountByUser(r.accountNbr, &r, u);

                printf("%s", u.name);
                printf("%lf\n", r.amount);
                if (r.amount < account_amount)
                {
                    printf("\nInsufficient funds");
                    stayOrReturn(1, u);
                    return;
                }
                r.amount = r.amount - account_amount;
                printf("%lf", r.amount);
                cr.amount = cr.amount + account_amount;
                fclose(pff);
                memcpy(u.name, TempUser.name, 32);
                break;
            }
        }
        if (CheckIfAccountIsOwnedByUser(u, u.name, cr.accountNbr) == 0 || CheckIfAccountIsOwnedByUser(u, u.name, r.accountNbr) == 0)
        {
            printf("\nYou do not own this account\n");
            stayOrReturn(1, u);
            return;
        }
        RemoveAccount(u, r.accountNbr);
        RemoveAccount(u, cr.accountNbr);
        FILE *writer = fopen(RECORDS, "a+");
        saveAccountToFile(writer, &u, &r);
        saveAccountToFile(writer, &u, &cr);
        fclose(writer);
        printf("%d", cr.accountNbr);
        success(u);
    }
}

void GetAccountByUser(int account_number, struct Record *result, struct User u)
{
    char userName[100];
    struct Record r;
    struct Record cr;
    struct User TempUser;
    struct User Temp2User;
    FILE *pff = fopen(RECORDS, "r+");

    system("clear");
    printf("\t\t\t===== New record =====\n");
    while (getAccountFromFile(pff, userName, &cr))
    {
        if (strcmp(userName, u.name) == 0)
        {
            if (cr.accountNbr == account_number)
            {
                memcpy(result, &cr, sizeof(cr));
                break;
            }
        }
    }
    fclose(pff);
}
