package models

type Permission string

const (
	PermissionCanDeleteUser                         Permission = "btcpay.user.candeleteuser"
	PermissionCanViewProfile                        Permission = "btcpay.user.canviewprofile"
	PermissionCanModifyProfile                      Permission = "btcpay.user.canmodifyprofile"
	PermissionCanManageNotificationsForUser         Permission = "btcpay.user.canmanagenotificationsforuser"
	PermissionCanViewNotificationsForUser           Permission = "btcpay.user.canviewnotificationsforuser"
	PermissionCanViewUsers                          Permission = "btcpay.server.canviewusers"
	PermissionCanCreateUser                         Permission = "btcpay.server.cancreateuser"
	PermissionCanManageUsers                        Permission = "btcpay.server.canmanageusers"
	PermissionCanModifyServerSettings               Permission = "btcpay.server.canmodifyserversettings"
	PermissionCanUseInternalLightningNode           Permission = "btcpay.server.canuseinternallightningnode"
	PermissionCanViewLightningInvoiceInternalNode   Permission = "btcpay.server.canviewlightninginvoiceinternalnode"
	PermissionCanCreateLightningInvoiceInternalNode Permission = "btcpay.server.cancreatelightninginvoiceinternalnode"
	PermissionCanModifyStoreSettings                Permission = "btcpay.store.canmodifystoresettings"
	PermissionCanViewCustodianAccounts              Permission = "btcpay.store.canviewcustodianaccounts"
	PermissionCanManageCustodianAccounts            Permission = "btcpay.store.canmanagecustodianaccounts"
	PermissionCanDepositToCustodianAccount          Permission = "btcpay.store.candeposittocustodianaccount"
	PermissionCanWithdrawDromCustodianAccount       Permission = "btcpay.store.canwithdrawfromcustodianaccount"
	PermissionCanTradeCustodianAccount              Permission = "btcpay.store.cantradecustodianaccount"
	PermissionWebhooks                              Permission = "btcpay.store.webhooks"
	PermissionCanViewStoreSettings                  Permission = "btcpay.store.canviewstoresettings"
	PermissionCanCreateInvoice                      Permission = "btcpay.store.cancreateinvoice"
	PermissionCanViewInvoices                       Permission = "btcpay.store.canviewinvoices"
	PermissionCanModifyInvoices                     Permission = "btcpay.store.canmodifyinvoices"
	PermissionCanModifyPaymentRequests              Permission = "btcpay.store.canmodifypaymentrequests"
	PermissionCanViewPaymentRequests                Permission = "btcpay.store.canviewpaymentrequests"
	PermissionCanManagePullPayments                 Permission = "btcpay.store.canmanagepullpayments"
	PermissionCanCreatePullPayments                 Permission = "btcpay.store.cancreatepullpayments"
	PermissionCanCreateNonApprovedPullPayments      Permission = "btcpay.store.cancreatenonapprovedpullpayments"
	PermissionCanUseLightningNode                   Permission = "btcpay.store.canuselightningnode"
	PermissionCanViewLightningInvoice               Permission = "btcpay.store.canviewlightninginvoice"
	PermissionCanCreateLightningInvoice             Permission = "btcpay.store.cancreatelightninginvoice"
)
