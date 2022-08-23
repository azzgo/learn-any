#!/bin/env perl
use strict;
use warnings;


my @hars = `ls ./raw-har/*.har`;

if (not -d "dist") {
  mkdir "./dist" or die "can not create folder dist";
}

foreach my $har (@hars) {
  $har =~ s/^\s+|\s+$//g;
  my $jsonEntries = qx(
    cat $har | jq ".log.entries | map(select(.response.status == 200 and .response.content.mimeType == \\"application\/json\\"))"
  );

  my $har_file_name = `basename $har`;


  my $wfile = undef;
  open($wfile, ">dist/$har_file_name") or die "can not write to file dist/$har_file_name";
  print $wfile $jsonEntries;
  close $wfile;
}


