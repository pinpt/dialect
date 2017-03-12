package objc

import (
	"strings"
	"testing"

	"github.com/pinpt/dialect"
)

func TestObjectiveC(t *testing.T) {
	reader := strings.NewReader(`import <Foundation/Foundation.h>
@interface Foo : NSObject
@end

@implementation Foo
@end
`)
	result, err := dialect.Examine("Objective-C", "foo.m", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 7 {
		t.Fatalf("result.Loc should have been 7, was %d", result.Loc)
	}
	if result.Sloc != 5 {
		t.Fatalf("result.Sloc should have been 5, was %d", result.Sloc)
	}
	if result.Comments != 0 {
		t.Fatalf("result.Comments should have been 0, was %d", result.Comments)
	}
	if result.Blanks != 2 {
		t.Fatalf("result.Blanks should have been 2, was %d", result.Blanks)
	}
	if result.IsTest {
		t.Fatal("result.IsTest should have been false, was true")
	}
}

func TestObjectiveCTestCase(t *testing.T) {
	reader := strings.NewReader(`#import <XCTest/XCTest.h>

@interface SampleCalcTests : XCTestCase
@end

@implementation SampleCalcTests

- (void)setUp {
    [super setUp];
    // Put setup code here. This method is called before the invocation of each test method in the class.
}

- (void)tearDown {
    // Put teardown code here. This method is called after the invocation of each test method in the class.
    [super tearDown];
}

- (void)testExample {
    // This is an example of a functional test case.
    // Use XCTAssert and related functions to verify your tests produce the correct results.
}

- (void)testPerformanceExample {
    // This is an example of a performance test case.
    [self measureBlock:^{
        // Put the code you want to measure the time of here.
    }];
}
@end
`)
	result, err := dialect.Examine("Objective-C", "foo.m", reader, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("result was nil")
	}
	if result.Loc != 30 {
		t.Fatalf("result.Loc should have been 30, was %d", result.Loc)
	}
	if result.Sloc != 17 {
		t.Fatalf("result.Sloc should have been 17, was %d", result.Sloc)
	}
	if result.Comments != 6 {
		t.Fatalf("result.Comments should have been 6, was %d", result.Comments)
	}
	if result.Blanks != 7 {
		t.Fatalf("result.Blanks should have been 7, was %d", result.Blanks)
	}
	if result.IsTest == false {
		t.Fatal("result.IsTest should have been true, was false")
	}
}
