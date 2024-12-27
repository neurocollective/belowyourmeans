just the beginning

### Auth

Using a `map[string]string` for now, will be redis later.

`cookie -> userId`

cookie is stored in browser as `BelowYourMeans=$cookieToken`

#### most common us banks for checking accounts

Wells Fargo - NO, they are rotten to the core, wake up Warren Buffett

Goldman Sachs - NO, rotten to the core since the 1920s at least

Citibank - NO under consent decree, involved in 2008 crisis
https://en.wikipedia.org/wiki/Citibank#Controversies
https://en.wikipedia.org/wiki/Citigroup#Regulatory_action,_lawsuits,_and_arbitration
https://en.wikipedia.org/wiki/Citigroup#Criticism

Bank of America Financial Center - NO, TOO BIG TO FAIL
https://en.wikipedia.org/wiki/Bank_of_America#Lawsuits,_controversies,_and_incidents

Chase - NO, TOO BIG TO FAIL
https://en.wikipedia.org/wiki/JPMorgan_Chase#Lawsuits_and_legal_settlements_by_years

Capital One - Not clean but compare their record to the "NO"s above
https://en.wikipedia.org/wiki/Capital_One#Investigations_and_legal_actions

PNC
U.S. Bank
BMO
Ally
CIT
Citizens Bank
TD Bank
Truist Bank
Huntington bank
KeyBank

## pg_dump notes

```
# -d dbname
# --dbname=dbname
# Specifies the name of the database to connect to. This is equivalent to specifying dbname as the first non-option argument on the command line. The dbname
# can be a connection string. If so, connection string parameters will override any conflicting command line options.

# -h host
# --host=host
# Specifies the host name of the machine on which the server is running. If the value begins with a slash, it is used as the directory for the Unix domain
# socket. The default is taken from the PGHOST environment variable, if set, else a Unix domain socket connection is attempted.

# -p port
# --port=port
# Specifies the TCP port or local Unix domain socket file extension on which the server is listening for connections. Defaults to the PGPORT environment
# variable, if set, or a compiled-in default.

# -U username
# --username=username
# User name to connect as.

# -f file
# --file=file
# Send output to the specified file. This parameter can be omitted for file based output formats, in which case the standard output is used. It must be given
# for the directory output format however, where it specifies the target directory instead of a file. In this case the directory is created by pg_dump and must
# not exist before.
```