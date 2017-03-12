package csharp

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestCSharp(t *testing.T) {
	reader := strings.NewReader(`/**
 * Hello
 */
using System;

namespace Test {
    public class MyClass {
    }
 }
`)
	result, err := dialect.Examine("C#", "foo.cs", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 10 {
		t.Fatalf("result.Loc should have been 10, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestCSharpSelenium(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/C%23+Test+Setup+Example
	reader := strings.NewReader(`using NUnit.Framework;
using System;
using Selenium;
using System.Web;
using System.Text;
using System.Net;
using OpenQA.Selenium;
using OpenQA.Selenium.Remote;
using OpenQA.Selenium.Support.UI;

namespace Saucey_Selenium {
    [TestFixture("chrome", "45", "Windows 7", "", "")]
    public class SauceNUnit_Test
    {
        private IWebDriver driver;
        private String browser;
        private String version;
        private String os;
        private String deviceName;
        private String deviceOrientation;

        public SauceNUnit_Test(String browser, String version, String os, String deviceName, String deviceOrientation)
        {
            this.browser = browser;
            this.version = version;
            this.os = os;
            this.deviceName = deviceName;
            this.deviceOrientation = deviceOrientation;
        }

        [SetUp]
        public void Init()
        {
            DesiredCapabilities caps = new DesiredCapabilities();
            caps.SetCapability(CapabilityType.BrowserName, browser);
            caps.SetCapability(CapabilityType.Version, version);
            caps.SetCapability(CapabilityType.Platform, os);
            caps.SetCapability("deviceName", deviceName);
            caps.SetCapability("deviceOrientation", deviceOrientation);
            caps.SetCapability("username", "SAUCE_USERNAME");
            caps.SetCapability("accessKey", "SAUCE_ACCESS_KEY");
            caps.SetCapability("name", TestContext.CurrentContext.Test.Name);

            driver = new RemoteWebDriver(new Uri("http://ondemand.saucelabs.com:80/wd/hub"), caps, TimeSpan.FromSeconds(600))

        }

        [Test]
        public void googleTest()
        {
            driver.Navigate().GoToUrl("http://www.google.com");
            StringAssert.Contains("Google", driver.Title);
            IWebElement query = driver.FindElement(By.Name("q"));
            query.SendKeys("Sauce Labs");
            query.Submit();
        }

        [TearDown]
        public void CleanUp()
        {
            bool passed = TestContext.CurrentContext.Result.Status == TestStatus.Passed;
            try
            {
                // Logs the result to Sauce Labs
                ((IJavaScriptExecutor)driver).ExecuteScript("sauce:job-result=" + (passed ? "passed" : "failed"));
            }
            finally
            {
                // Terminates the remote webdriver session
                driver.Quit();
            }
        }
    }
}
`)
	result, err := dialect.Examine("C#", "foo.cs", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 75 {
		t.Fatalf("result.Loc should have been 75, was %d", result.Loc)
	}
	if result.Sloc != 65 {
		t.Fatalf("result.Sloc should have been 65, was %d", result.Sloc)
	}
	if result.Comments != 2 {
		t.Fatalf("result.Comments should have been 2, was %d", result.Comments)
	}
	if result.Blanks != 8 {
		t.Fatalf("result.Blanks should have been 8, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
