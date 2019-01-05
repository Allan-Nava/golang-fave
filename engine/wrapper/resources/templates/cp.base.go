package templates

var CpBase = []byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="theme-color" content="#205081" /><meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"><title>{{$.Data.Title}}</title><link rel="stylesheet" href="{{$.System.PathCssBootstrap}}"><link rel="stylesheet" href="{{$.System.PathCssStyles}}" /><link rel="stylesheet" href="{{$.System.PathCssCpStyles}}"><link rel="shortcut icon" href="{{$.System.PathIcoFav}}" type="image/x-icon" /></head><body class="cp cp-sidebar-left cp-sidebar-right"><div class="modal fade" id="sys-modal-user-settings" tabindex="-1" role="dialog" aria-labelledby="sysModalUserSettingsLabel" aria-hidden="true"><div class="modal-dialog modal-dialog-centered" role="document"><div class="modal-content"><form class="form-user-settings" action="/cp/" method="post" autocomplete="off"><input type="hidden" name="action" value="usersettings"><div class="modal-header"><h5 class="modal-title" id="sysModalUserSettingsLabel">Settings</h5><button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button></div><div class="modal-body text-left"><div class="form-group"><label for="first_name">User first name</label><input type="text" class="form-control" id="first_name" name="first_name" value="{{$.Data.UserFirstName}}" placeholder="User first name" autocomplete="off"></div><div class="form-group"><label for="last_name">User last name</label><input type="text" class="form-control" id="last_name" name="last_name" value="{{$.Data.UserLastName}}" placeholder="User last name" autocomplete="off"></div><div class="form-group"><label for="email">User email</label><input type="email" class="form-control" id="email" name="email" value="{{$.Data.UserEmail}}" placeholder="User email" autocomplete="off" required></div><div class="form-group"><label for="password">User new password</label><input type="password" class="form-control" id="password" name="password" value="{{$.Data.UserPassword}}" placeholder="User new password" autocomplete="off"></div></div><div class="modal-footer"><button type="submit" class="btn btn-primary">Save changes</button><button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button></div></form></div></div></div><nav class="navbar main navbar-expand-md navbar-dark fixed-top bg-dark"><a class="navbar-brand" href="/">{{$.Data.Title}}</a><button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button><div class="collapse navbar-collapse" id="navbarCollapse"><ul class="navbar-nav mr-auto"><li class="nav-item dropdown"><a class="nav-link dropdown-toggle" href="javascript:;" id="nbModulesDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Modules</a><div class="dropdown-menu" aria-labelledby="nbModulesDropdown"><a class="dropdown-item active" href="/cp/">Pages</a></div></li><li class="nav-item dropdown"><a class="nav-link dropdown-toggle" href="javascript:;" id="nbSystemDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">System</a><div class="dropdown-menu" aria-labelledby="nbSystemDropdown"><a class="dropdown-item" href="/cp/users/">Users</a><div class="dropdown-divider"></div><a class="dropdown-item" href="/cp/settings/">Settings</a></div></li></ul><ul class="navbar-nav ml-auto"><li class="nav-item dropdown"><a class="nav-link dropdown-toggle" href="javascript:;" id="nbAccountDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><img class="rounded-circle" src="{{$.Data.UserAvatarLink}}">{{$.Data.UserEmail}}</a><div class="dropdown-menu dropdown-menu-right" aria-labelledby="nbAccountDropdown"><a class="dropdown-item" href="javascript:ActionUserSettings();" data-toggle="modal" data-target="#sys-modal-user-settings">Settings</a><div class="dropdown-divider"></div><a class="dropdown-item" href="javascript:ActionSingOut();">Sing out</a></div></li></ul></div></nav><div class="wrap"><div class="sidebar sidebar-left"><div class="scroll"><div class="padd">{{$.Data.SidebarLeft}}</div></div></div><div class="content"><div class="scroll"><div class="padd">{{$.Data.Content}}</div></div></div><div class="sidebar sidebar-right"><div class="scroll"><div class="padd">{{$.Data.SidebarRight}}</div></div></div></div><script src="{{$.System.PathJsJquery}}"></script><script src="{{$.System.PathJsPopper}}"></script><script src="{{$.System.PathJsBootstrap}}"></script><script src="{{$.System.PathJsCpScripts}}"></script></body></html>`)
