package java

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
	_ "github.com/pinpt/dialect/pkg/java"
)

func TestJava(t *testing.T) {
	reader := strings.NewReader(`/**
 * Hello
 */
package foo

public class Test {

}
`)
	result, err := dialect.Examine("Java", "foo.java", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 9 {
		t.Fatalf("result.Loc should have been 9, was %d", result.Loc)
	}
	if result.Sloc != 3 {
		t.Fatalf("result.Sloc should have been 3, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 3 {
		t.Fatalf("result.Blanks should have been 3, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestJavaSelenium(t *testing.T) {
	// from https://wiki.saucelabs.com/display/DOCS/Java+Test+Setup+Example#JavaTestSetupExample-CodeExample
	reader := strings.NewReader(`import org.openqa.selenium.WebDriver;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.remote.RemoteWebDriver;

import java.net.URL;

public class SampleSauceTest {

	public static final String USERNAME = "YOUR_USERNAME";
	public static final String ACCESS_KEY = "YOUR_ACCESS_KEY";
	public static final String URL = "https://" + USERNAME + ":" + ACCESS_KEY + "@ondemand.saucelabs.com:443/wd/hub";

	public static void main(String[] args) throws Exception {

	 DesiredCapabilities caps = DesiredCapabilities.chrome();
	 caps.setCapability("platform", "Windows XP");
	 caps.setCapability("version", "43.0");

	 WebDriver driver = new RemoteWebDriver(new URL(URL), caps);

	 /**
	  * Goes to Sauce Lab's guinea-pig page and prints title
	  */

	 driver.get("https://saucelabs.com/test/guinea-pig");
	 System.out.println("title of page is: " + driver.getTitle());

	 driver.quit();
	}
}
`)
	result, err := dialect.Examine("Java", "foo.java", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 31 {
		t.Fatalf("result.Loc should have been 31, was %d", result.Loc)
	}
	if result.Sloc != 18 {
		t.Fatalf("result.Sloc should have been 18, was %d", result.Sloc)
	}
	if result.Comments != 3 {
		t.Fatalf("result.Comments should have been 3, was %d", result.Comments)
	}
	if result.Blanks != 10 {
		t.Fatalf("result.Blanks should have been 10, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}

func TestJavaJUnit(t *testing.T) {
	// from https://www.tutorialspoint.com/junit/junit_test_framework.htm
	reader := strings.NewReader(`import junit.framework.*;

public class JavaTest extends TestCase {
   protected int value1, value2;

   // assigning the values
   protected void setUp(){
      value1 = 3;
      value2 = 3;
   }

   // test method to add two values
   public void testAdd(){
      double result = value1 + value2;
      assertTrue(result == 6);
   }
}
`)
	result, err := dialect.Examine("Java", "foo.java", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 18 {
		t.Fatalf("result.Loc should have been 18, was %d", result.Loc)
	}
	if result.Sloc != 12 {
		t.Fatalf("result.Sloc should have been 12, was %d", result.Sloc)
	}
	if result.Comments != 2 {
		t.Fatalf("result.Comments should have been 2, was %d", result.Comments)
	}
	if result.Blanks != 4 {
		t.Fatalf("result.Blanks should have been 4, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
