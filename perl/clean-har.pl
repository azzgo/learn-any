#!/bin/env perl
use strict;
use warnings;


my @hars = `ls ./raw-har/*.har`;

if (not -d "dist") {
  mkdir "./dist" or die "can not create folder dist";
}

foreach my $har (@hars) {
  $har =~ s/^\s+|\s+$//g;
  my $url = qx(
    cat $har | jq ".log.entries[0].request.url"
  );

  my $har_file_name = `basename $har`;


  my $wfile = undef;
  open($wfile, ">dist/$har_file_name") or die "can not write to file dist/$har_file_name";
  print $wfile $url;
  close $wfile;
}


