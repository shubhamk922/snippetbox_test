package doc

//It’s important to point out that the directory name internal carries a special meaning and behavior in Go: any packages which live under this directory can only be imported by code inside the parent of the internal directory. In our case, this means that any packages which live in internal can only be imported by code inside our snippetbox project directory.
