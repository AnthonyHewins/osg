package investigate

import "testing"

func TestContainsUrl(t *testing.T) {
	for _, url_string := range is_a_url_and_so_it_should_be_flagged {
		if contains_url(&url_string) != 0 { continue }
		t.Errorf("should be classified as URL: %v", url_string)
	}

	for _, string_no_url := range not_a_url_so_shouldnt_be_suspicious {
		if contains_url(&string_no_url) == 0 { continue }
		t.Errorf("should NOT be classified as URL: %v", string_no_url)
	}
}

func TestContainsSuspiciousWord(t *testing.T) {
	for _, suspicious_filename := range suspicious_filenames {
		if contains_suspicious_word(&suspicious_filename) != 0 { continue }
		t.Errorf("should be classified as suspicious: %v", suspicious_filename)
	}

	for _, normal_filename := range not_suspicious_filenames {
		if contains_suspicious_word(&normal_filename) == 0 { continue }
		t.Errorf("should NOT be classified as suspicious: %v", normal_filename)
	}
}

var not_a_url_so_shouldnt_be_suspicious = []string{
	"about: Create a report to help us improve",
	"Hey there! Thanks for using Firefox!",
	"- [ ] Submit build to Apple (Select YES to IDFA, 'Attribute this app installation to a previously served advertisement')",
	"- variable_name",
	"- type_name",
	"- cyclomatic_complexity",
	"- trailing_comma",
	"- discarded_notification_center_observer",
	"- switch_case_alignment",
	"- duplicate_imports",
	"- overridden_super_call",
	"- prohibited_super_call",
	"- Client/Assets/Search/get_supported_locales.swift",
	"# configurable rules can be customized from this configuration file",
	"# binary rules can set their severity level",
	"force_cast: warning",
	"here, you may add your name and, optionally, email address in the",
	"/// The auth endpoint user identifier identifying the account.  A Firefox Account is uniquely identified on a",
	"// Use updateDeviceName() to set the device name and update the device registration.",
	"private(set) open var commandsClient: FxACommandsClient!",
	"// Add notification on state change (checkpoint is called on each change)",
	"NotificationCenter.default.post(name: .FirefoxAccountStateChange, object: nil)",
	"// advance() is in progress, the shared deferred will be returned.  (Multiple consumers can chain off a single",
	"// deferred safely.)  If no advance() is in progress, a new shared deferred will be scheduled and returned.  To",
	"// As of this writing, the current version, v2, is backward compatible with v1. The only",
	"account.pushRegistration = dictionary[pushRegistration] as? PushRegistration",
	"self.image = UIImage(named: placeholder-avatar)",
	"NotificationCenter.default.post(name: .FirefoxAccountProfileChanged, object: self)",
	"// Don't forget to call Profile.flushAccount() to persist this change!",
	"// Fetch current user's FxA profile. It contains the most updated email, displayName and avatar. This",
	"// emits two `NotificationFirefoxAccountProfileChanged`, once when the profile has been downloaded and",
	"NotificationCenter.default.post(name: .FirefoxAccountProfileChanged, object: self)",
	"// If we have a cached copy of the KeyID in the Keychain, use it.",
	"if let cachedOAuthKeyID = KeychainStore.shared.string(forKey: kidKeychainKey) {",
	"return deferMaybe(cachedOAuthKeyID)",
	"return deferMaybe(FxAClientError.local(NSError()))",
	"return deferMaybe(FxAClientError.local(NSError()))",
	"// Clear cached FxA profile data from Keychain.",
	"// Clear cached OAuth data from Keychain.",
	"NotificationCenter.default.post(name: .FirefoxAccountDeviceRegistrationUpdated, object: nil)",
	"return JSON(commands)",
	"// We already have an advance() in progress.  This consumer can chain from it.",
	"let cachedState = stateCache.value!",
	"// If we were not previously married, but we are now,",
	"@discardableResult open func makeDoghouse() -> Bool {",
	"case .custom: return CustomFirefoxAccountConfiguration(prefs: prefs)",
	"* and context=fx_ios_v1 opts us in to the Desktop Sync postMessage interface.",
}

var is_a_url_and_so_it_should_be_flagged = []string{
	"If you are having issues with using Firefox make sure to check https://support.mozilla.org/en-US/products/ios first",
	"See [Release build checklist wiki](https://github.com/mozilla-mobile/firefox-ios/wiki/Release-Build-Checklist) for more detailed instructions.",
	"required_string: /* This Source Code Form is subject to the terms of the Mozilla Public\n * License, v. 2.0. If a copy of the MPL was not distributed with this\n * file, You can obtain one at http://mozilla.org/MPL/2.0/. */",
	"contribution to Mozilla, see http://www.mozilla.org/credits/.",
	"* file, You can obtain one at http://mozilla.org/MPL/2.0/. */",
	"/// https://github.com/mozilla/fxa-auth-server/blob/02f88502700b0c5ef5a4768a8adf332f062ad9bf/docs/api.md",
	"/// https://github.com/mozilla/fxa-oauth-server/blob/6cc91e285fc51045a365dbacb3617ef29093dbc3/docs/api.md",
	"* file, You can obtain one at http://mozilla.org/MPL/2.0/. */",
	"// From https://accounts.firefox.com/.well-known/fxa-client-configuration",
	"private let ProductionAuthEndpointURL = https://api.accounts.firefox.com/v1)!",
	"private let ProductionOAuthEndpointURL = https://oauth.accounts.firefox.com/v1)!",
	"private let ProductionProfileEndpointURL = https://profile.accounts.firefox.com/v1)!",
	"private let ProductionTokenServerEndpointURL = https://token.services.mozilla.com/1.0/sync/1.5)!",
	"private let ProductionSignInURL = https://accounts.firefox.com/signin?service=sync&context=fx_ios_v1)!",
	"private let ProductionSettingsURL = https://accounts.firefox.com/settings?context=fx_ios_v1)!",
	"private let ProductionForceAuthURL = https://accounts.firefox.com/force_auth?service=sync&context=fx_ios_v1)!",
	"// From https://accounts.firefox.com.cn/.well-known/fxa-client-configuration",
	"private let ChinaAuthEndpointURL = https://api-accounts.firefox.com.cn/v1)!",
	"private let ChinaOAuthEndpointURL = https://oauth.firefox.com.cn/v1)!",
	"private let ChinaProfileEndpointURL = https://profile.firefox.com.cn/v1)!",
	"private let ChinaTokenServerEndpointURL = https://sync.firefox.com.cn/token/1.0/sync/1.5)!",
	"private let ChinaSignInURL = https://accounts.firefox.com.cn/signin?service=sync&context=fx_ios_v1)!",
	"private let ChinaSettingsURL = https://accounts.firefox.com.cn/settings?context=fx_ios_v1)!",
	"private let ChinaForceAuthURL = https://accounts.firefox.com.cn/force_auth?service=sync&context=fx_ios_v1)!",
	"// From https://accounts.stage.mozaws.net/.well-known/fxa-client-configuration",
	"private let StageAuthEndpointURL = https://api-accounts.stage.mozaws.net/v1)!",
	"private let StageOAuthEndpointURL = https://oauth.stage.mozaws.net/v1)!",
	"private let StageProfileEndpointURL = https://profile.stage.mozaws.net/v1)!",
	"private let StageTokenServerEndpointURL = https://token.stage.mozaws.net/1.0/sync/1.5)!",
	"private let StageSignInURL = https://accounts.stage.mozaws.net/signin?service=sync&context=fx_ios_v1)!",
	"private let StageSettingsURL = https://accounts.stage.mozaws.net/settings?context=fx_ios_v1)!",
	"private let StageForceAuthURL = https://accounts.stage.mozaws.net/force_auth?service=sync&context=fx_ios_v1)!",
	"// From https://latest.dev.lcip.org/.well-known/fxa-client-configuration",
	"private let LatestDevAuthEndpointURL = https://latest.dev.lcip.org/auth/v1)!",
	"private let LatestDevOAuthEndpointURL = https://oauth-latest.dev.lcip.org)!",
	"private let LatestDevProfileEndpointURL = https://latest.dev.lcip.org/profile)!",
	"private let LatestDevSignInURL = https://latest.dev.lcip.org/signin?service=sync&context=fx_ios_v1)!",
	"private let LatestDevSettingsURL = https://latest.dev.lcip.org/settings?context=fx_ios_v1)!",
	"private let LatestDevForceAuthURL = https://latest.dev.lcip.org/force_auth?service=sync&context=fx_ios_v1)!",
	"// From https://stable.dev.lcip.org/.well-known/fxa-client-configuration",
	"public let StableDevAuthEndpointURL = https://stable.dev.lcip.org/auth/v1)!",
	"public let StableDevOAuthEndpointURL = https://oauth-stable.dev.lcip.org)!",
	"public let StableDevProfileEndpointURL = https://stable.dev.lcip.org/profile)!",
	"public let StableDevSignInURL = https://stable.dev.lcip.org/signin?service=sync&context=fx_ios_v1)!",
	"public let StableDevSettingsURL = https://stable.dev.lcip.org/settings?context=fx_ios_v1)!",
	"public let StableDevForceAuthURL = https://stable.dev.lcip.org/force_auth?service=sync&context=fx_ios_v1)!",
}

var not_suspicious_filenames = []string{
	"./buddybuild_carthage_command.sh",
	"./bootstrap.sh",
	"./buddybuild_postclone.sh",
	"./SyncTelemetry",
	"./SyncTelemetry/SyncTelemetryEvents.swift",
	"./SyncTelemetry/SyncTelemetry.swift",
	"./SyncTelemetry/SyncTelemetry.h",
	"./SyncTelemetry/Info.plist",
	"./l10n-screenshots.sh",
	"./SyncTelemetryTests",
	"./SyncTelemetryTests/EventTests.swift",
	"./SyncTelemetryTests/Info.plist",
	"./Providers",
	"./Providers/SyncStatusResolver.swift",
	"./Providers/Profile.swift",
	"./Providers/PocketFeed.swift",
	"./Storage",
	"./Storage/Visit.swift",
	"./Storage/Sharing.swift",
	"./Storage/ExtensionUtils.swift",
	"./Storage/ReadingList.swift",
	"./Storage/PageMetadata.swift",
	"./Storage/History.swift",
	"./Storage/DatabaseError.swift",
	"./Storage/Cursor.swift",
	"./Storage/Syncable.swift",
	"./Storage/Rust",
	"./Storage/Rust/RustPlaces.swift",
	"./Storage/Rust/RustLogins.swift",
	"./Storage/Rust/RustShared.swift",
	"./Storage/RecentlyClosedTabs.swift",
	"./Storage/Favicons.swift",
	"./Storage/Clients.swift",
	"./Storage/SyncQueue.swift",
	"./Storage/FileAccessor.swift",
	"./Storage/Metadata.swift",
	"./Storage/Queue.swift",
	"./Storage/DiskImageStore.swift",
	"./Storage/SuggestedSites.swift",
	"./Storage/ThirdParty",
	"./Storage/ThirdParty/SwiftData.swift",
	"./Storage/CertStore.swift",
	"./Storage/Storage-Bridging-Header.h",
	"./Storage/SQL",
	"./Storage/SQL/SQLiteHistory.swift",
	"./Storage/SQL/SQLiteMetadata.swift",
	"./Storage/SQL/SQLiteQueue.swift",
	"./Storage/SQL/SQLiteHistoryFavicons.swift",
	"./Storage/SQL/Schema.swift",
	"./Storage/SQL/SQLiteHistoryFactories.swift",
	"./Storage/SQL/BrowserSchema.swift",
	"./Storage/SQL/SQLiteRemoteClientsAndTabs.swift",
	"./Storage/SQL/BrowserDB.swift",
	"./Storage/SQL/SQLiteReadingList.swift",
	"./Storage/SQL/ReadingListSchema.swift",
	"./Storage/SQL/SQLiteFavicons.swift",
	"./Storage/SQL/SQLiteHistoryRecommendations.swift",
	"./Storage/Info.plist",
	"./Storage/RemoteTabs.swift",
	"./Storage/Site.swift",
	"./Storage/DefaultSuggestedSites.swift",
	"./MarketingUITests",
	"./MarketingUITests/MarketingUITests.swift",
	"./MarketingUITests/Info.plist",
	"./SharedTests",
	"./SharedTests/AsyncReducerTests.swift",
	"./SharedTests/SupportUtilsTests.swift",
	"./SharedTests/DeferredTestUtils.swift",
	"./SharedTests/RollingFileLoggerTests.swift",
	"./SharedTests/ArrayExtensionTests.swift",
	"./SharedTests/FeatureSwitchTests.swift",
	"./SharedTests/AuthenticationKeychainInfoTests.swift",
	"./SharedTests/UtilsTests.swift",
	"./SharedTests/HexExtensionsTests.swift",
	"./SharedTests/ResultTests.swift",
	"./SharedTests/DeferredTests.swift",
	"./SharedTests/Info.plist",
	"./SharedTests/NSURLExtensionsTests.swift",
	"./Dangerfile",
	"./AUTHORS",
	"./content-blocker-lib-ios",
	"./content-blocker-lib-ios/base-fingerprinting-track.json",
	"./content-blocker-lib-ios/ContentBlockerGen",
	"./content-blocker-lib-ios/ContentBlockerGen/Package.swift",
	"./content-blocker-lib-ios/ContentBlockerGen/Sources",
	"./content-blocker-lib-ios/ContentBlockerGen/Sources/ContentBlockerGen",
	"./content-blocker-lib-ios/ContentBlockerGen/Sources/ContentBlockerGen/main.swift",
	"./content-blocker-lib-ios/ContentBlockerGen/Sources/ContentBlockerGenLib",
	"./content-blocker-lib-ios/ContentBlockerGen/Sources/ContentBlockerGenLib/ContentBlockerGenLib.swift",
	"./content-blocker-lib-ios/ContentBlockerGen/Tests",
	"./content-blocker-lib-ios/ContentBlockerGen/Tests/ContentBlockerGenTests",
	"./content-blocker-lib-ios/ContentBlockerGen/Tests/ContentBlockerGenTests/XCTestManifests.swift",
	"./content-blocker-lib-ios/ContentBlockerGen/Tests/ContentBlockerGenTests/ContentBlockerGenTests.swift",
	"./content-blocker-lib-ios/js",
	"./content-blocker-lib-ios/js/TrackingProtectionStats.js",
	"./content-blocker-lib-ios/src",
	"./content-blocker-lib-ios/src/ContentBlocker.swift",
	"./content-blocker-lib-ios/src/TabContentBlocker.swift",
	"./content-blocker-lib-ios/src/TrackingProtectionPageStats.swift",
}

var suspicious_filenames = []string{
	"analyticsService",
	"analytic_docker.rs",
	"js_analytic_agent",
	"telemetry_service",
}