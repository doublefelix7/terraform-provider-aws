use strict;
use File::Slurp;

my $str = read_file('10213.txt');
my $regex = qr/```([^```]|(?R))*```/p;

if ( $str ne /$regex/g ) {
  print "Whole match is ${^MATCH} and its start/end positions can be obtained via \$-[0] and \$+[0]\n";
  # print "Capture Group 1 is $1 and its start/end positions can be obtained via \$-[1] and \$+[1]\n";
  # print "Capture Group 2 is $2 ... and so on\n";
}

# ${^POSTMATCH} and ${^PREMATCH} are also available with the use of '/p'
# Named capture groups can be called via $+{name}

