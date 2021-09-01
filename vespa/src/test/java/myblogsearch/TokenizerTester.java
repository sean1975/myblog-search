package myblogsearch;

import com.yahoo.language.Language;
import com.yahoo.language.Linguistics;
import com.yahoo.language.process.StemMode;
import com.yahoo.language.process.Token;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import myblogsearch.config.SimpleChineseLinguisticsConfig;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class TokenizerTester {

    private boolean accentDrop = false;
    private Language language = Language.ENGLISH;
    private StemMode stemMode = StemMode.NONE;

    public void assertTokens(String input, String ... expectedTokenStrings) {
        List<String> actual = new ArrayList<>();
        for (Token token : tokenize(input)) {
            findTokenStrings(token, actual);
        }
        assertEquals(Arrays.asList(expectedTokenStrings), actual);
    }

    public List<String> findTokenStrings(Token token, List<String> out) {
        int numComponents = token.getNumComponents();
        if (token.isSpecialToken() || numComponents == 0) {
            out.add(token.getTokenString());
        } else {
            for (int i = 0; i < numComponents; ++i) {
                findTokenStrings(token.getComponent(i), out);
            }
        }
        return out;
    }

    public Iterable<Token> tokenize(String input) {
        SimpleChineseLinguisticsConfig.Builder builder = new SimpleChineseLinguisticsConfig.Builder().dictionaryPath("/cities.txt");
        Linguistics linguistics = new SimpleChineseLinguistics(builder.build());
        return linguistics.getTokenizer().tokenize(input, language, stemMode, accentDrop);
    }

    public TokenizerTester setAccentDrop(boolean accentDrop) {
        this.accentDrop = accentDrop;
        return this;
    }

    public TokenizerTester setLanguage(Language language) {
        this.language = language;
        return this;
    }

    public TokenizerTester setStemMode(StemMode stemMode) {
        this.stemMode = stemMode;
        return this;
    }

}
