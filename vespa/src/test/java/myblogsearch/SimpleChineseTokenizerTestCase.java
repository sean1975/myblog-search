package myblogsearch;

import com.yahoo.language.Language;
import com.yahoo.language.process.StemMode;
import com.yahoo.language.process.Token;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

public class SimpleChineseTokenizerTestCase {

    @Test
    public void testTokenizing() {
        TokenizerTester tester = new TokenizerTester();
        assertNotNull(tester, "Failed to create TokenizerTester");
        tester.assertTokens(
            "However, my classic iPhone 4S is not compatible with the government app for scanning the QR code.",
            "however", ",", "my", "classic", "iphone", "4", "s", "is", "not", "compatible", "with", "the", "government", "app", "for", "scanning", "the", "qr", "code", "."
        );
    }

    @Test
    public void testTokenizingStemming() {
        TokenizerTester tester = new TokenizerTester().setStemMode(StemMode.ALL);
        tester.assertTokens(
            "However, my classic iPhone 4S is not compatible with the government app for scanning the QR code.",
            "howev", ",", "my", "classic", "iphon", "4", "s", "is", "not", "compat", "with", "the", "govern", "app", "for", "scan", "the", "qr", "code", "."
        );
    }

    @Test
    public void testTokenizingChinese() {
        TokenizerTester tester = new TokenizerTester().setStemMode(StemMode.ALL).setLanguage(Language.CHINESE_TRADITIONAL);
        tester.assertTokens(
            "很不幸地，我的經典款 iPhone 4S 不能安裝州政府的 app。",
            "不幸", "經典", "款", " ", "iphon", " ", "4s", " ", "不能", "安裝", "州政府", " ", "app"
        );
    }
}
